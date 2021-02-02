package gosatellite

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion  = "0.1.0"
	defaultBaseURL  = "https://satellite.example.com"
	userAgent       = "gosatellite/" + libraryVersion
	mediaType       = "application/json"
	basePath        = "/api"
	katelloBasePath = "/katello/api"
)

// Config defines the configuration needed to connect to the
// Red Hat Satellite server
type Config struct {
	Username      string
	Password      string
	SatelliteHost string
	SSLVerify     bool
}

// Client is the API client for Red Hat Satellite
type Client struct {
	// HTTP client used to communicate with the Red Hat Satellite API.
	client *http.Client

	// Config for the client such as the satellite hostname and credentials
	Config *Config

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	ActivationKeys  ActivationKeys
	AuthSourceLDAPs AuthSourceLDAPs
	Filters         Filters
	HostCollections HostCollections
	Locations       Locations
	Manifests       Manifests
	Organizations   Organizations
	Permissions     Permissions
	Products        Products
	Repositories    Repositories
	Roles           Roles
	UserGroups      UserGroups

	// Optional function called after every successful request made to the Red Hat Satellite APIs
	onRequestCompleted RequestCompletionCallback

	// Optional extra HTTP headers to set on every request to the API.
	headers map[string]string
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// Scope by locations
	LocationID int `url:"location_id,omitempty"`

	// Sort field and order, eg. ‘id DESC’
	Order string `url:"order,omitempty"`

	// Scope by organizations
	OrganizationID int `url:"organization_id,omitempty"`

	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`

	// A search string used to filter results
	Search string `url:"search,omitempty"`
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	ErrorStruct *struct {
		FullMessages *[]string `json:"full_messages"`
		Message      *string   `json:"message"`
	} `json:"error"`
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}

// NewClient returns a new Red Hat Satellite REST API client
func NewClient(config *Config) (*Client, error) {
	defaultBaseURL := "https://" + config.SatelliteHost
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	if !config.SSLVerify {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	c := &Client{client: http.DefaultClient, BaseURL: baseURL, UserAgent: userAgent, Config: config}
	c.ActivationKeys = &ActivationKeysOp{client: c}
	c.AuthSourceLDAPs = &AuthSourceLDAPsOp{client: c}
	c.Filters = &FiltersOp{client: c}
	c.HostCollections = &HostCollectionsOp{client: c}
	c.Locations = &LocationsOp{client: c}
	c.Manifests = &ManifestsOp{client: c}
	c.Organizations = &OrganizationsOp{client: c}
	c.Permissions = &PermissionsOp{client: c}
	c.Products = &ProductsOp{client: c}
	c.Repositories = &RepositoriesOp{client: c}
	c.Roles = &RolesOp{client: c}
	c.UserGroups = &UserGroupsOp{client: c}

	c.headers = make(map[string]string)

	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}

	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(method, u.String(), buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", mediaType)
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	req.SetBasicAuth(c.Config.Username, c.Config.Password)

	return req, nil
}

// NewManifestUploadRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. The content in form is
// added to the request as form data
func (c *Client) NewManifestUploadRequest(ctx context.Context, method, urlStr string, manifest []byte, manifestFilename string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	var fw io.Writer
	if fw, err = w.CreateFormFile("content", manifestFilename); err != nil {
		return nil, err
	}
	if _, err = fw.Write(manifest); err != nil {
		return nil, err
	}
	w.Close()

	req, err := http.NewRequest(method, u.String(), &buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Config.Username, c.Config.Password)

	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Header.Add("Multipart", "true")
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// OnRequestCompleted sets the Red Hat Satellite API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := DoRequestWithClient(ctx, c.client, req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}

// DoRequest submits an HTTP request.
func DoRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	return DoRequestWithClient(ctx, http.DefaultClient, req)
}

// DoRequestWithClient submits an HTTP request using the specified client.
func DoRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

func (r *ErrorResponse) Error() string {
	allMessages := []string{}
	if r.ErrorStruct != nil {
		if r.ErrorStruct.FullMessages != nil {
			allMessages = append(allMessages, *r.ErrorStruct.FullMessages...)
		}
		if r.ErrorStruct.Message != nil {
			allMessages = append(allMessages, *r.ErrorStruct.Message)
		}
	}
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, strings.Join(allMessages, "|"))
}

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.ErrorStruct.FullMessages = &[]string{string(data)}
		}
	}

	return errorResponse
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}
