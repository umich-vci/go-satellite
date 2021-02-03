package gosatellite

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

const activationKeyPath = katelloBasePath + "/activation_keys"

// ActivationKey defines model for an Activation Key.
type ActivationKey struct {
	AutoAttach       *bool                 `json:"auto_attach"`
	ContentOverrides *[]akContentOverrides `json:"content_overrides"`
	ContentView      *genericShortRef      `json:"content_view"`
	ContentViewID    *int                  `json:"content_view_id"`
	CreatedAt        *string               `json:"created_at"`
	Description      *string               `json:"description"`
	Environment      *genericShortRef      `json:"environment"`
	EnvironmentID    *int                  `json:"environment_id"`
	HostCollections  *[]genericShortRef    `json:"host_collections"`
	ID               *int                  `json:"id"`
	MaxHosts         *int                  `json:"max_hosts"`
	Name             *string               `json:"name"`
	Organization     *shortOrg             `json:"organization"`
	OrganizationID   *int                  `json:"organization_id"`
	Permissions      *akPermissions        `json:"permissions"`
	Products         *[]genericShortRef    `json:"products"`
	ReleaseVersion   *string               `json:"release_version"`
	ServiceLevel     *string               `json:"service_level"`
	UnlimitedHosts   *bool                 `json:"unlimited_hosts"`
	UpdatedAt        *string               `json:"updated_at"`
	UsageCount       *int                  `json:"usage_count"`
	UserID           *int                  `json:"user_id"`
}

// ActivationKeyList defines model for a list of activation keys.
type ActivationKeyList struct {
	searchResults
	Error   *string          `json:"error"`
	Results *[]ActivationKey `json:"results"`
}

// ActivationKeyReleasesList defines model for a list of releases available for an activation keys.
type ActivationKeyReleasesList struct {
	searchResults
	Error   *string   `json:"error"`
	Results *[]string `json:"results"`
}

// ActivationKeyProductContentList defines model for a list of products available for an activation keys.
type ActivationKeyProductContentList struct {
	searchResults
	Error   *string   `json:"error"`
	Results *[]string `json:"results"`
}

// ActivationKeyListOptions specifies the optional parameters to various List methods that
// support pagination.
type ActivationKeyListOptions struct {
	KatelloListOptions

	// Filter by content view id
	ContentViewID int `url:"content_view_id,omitempty"`

	// Scope by environment
	EnvironmentID int `url:"environment_id,omitempty"`

	// Filter by name
	Name string `url:"name,omitempty"`

	// Scope by organizations
	OrganizationID int `url:"organization_id,omitempty"`
}

// ActivationKeyAvailableHostCollectionsListOptions specifies the optional parameters to various List methods that
// support pagination.
type ActivationKeyAvailableHostCollectionsListOptions struct {
	KatelloListOptions

	// Filter by content view id
	ContentViewID int `url:"content_view_id,omitempty"`

	// Scope by environment
	EnvironmentID int `url:"environment_id,omitempty"`

	// Filter by name
	Name string `url:"name,omitempty"`

	// Scope by organizations
	OrganizationID int `url:"organization_id,omitempty"`
}

// ActivationKeyCreate defines model for creating an activation key.
type ActivationKeyCreate struct {
	OrganizationID *int    `json:"organization_id"`
	Name           *string `json:"name"`
	Description    *string `json:"description,omitempty"`
	EnvironmentID  *int    `json:"environment_id,omitempty"`
	ContentViewID  *int    `json:"content_view_id,omitempty"`
	MaxHosts       *int    `json:"max_hosts,omitempty"`
	UnlimitedHosts *bool   `json:"unlimited_hosts,omitempty"`
}

// ActivationKeyUpdate defines model for updating an activation key.
type ActivationKeyUpdate struct {
	OrganizationID *int    `json:"organization_id,omitempty"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	EnvironmentID  *int    `json:"environment_id,omitempty"`
	ContentViewID  *int    `json:"content_view_id,omitempty"`
	MaxHosts       *int    `json:"max_hosts,omitempty"`
	UnlimitedHosts *bool   `json:"unlimited_hosts,omitempty"`
	ReleaseVersion *string `json:"release_version,omitempty"`
	ServiceLevel   *string `json:"service_level,omitempty"`
	AutoAttach     *bool   `json:"auto_attach,omitempty"`
}

// ActivationKeyContentOverride defines model for defining a content override for an activation key.
type ActivationKeyContentOverride struct {
	ContentOverrides struct {
		ContentLabel *string `json:"content_label"`
		Value        *string `json:"value,omitempty"`
		Name         *string `json:"name,omitempty"`
		Remove       *bool   `json:"remove,omitempty"`
	} `json:"content_overrides"`
}

type akContentOverrides struct {
	ContentLabel *string `json:"content_label"`
	Name         *string `json:"name"`
	Value        *string `json:"value"`
}

type akPermissions struct {
	DestroyActivationKeys *bool `json:"destroy_activation_keys"`
	EditActivationKeys    *bool `json:"edit_activation_keys"`
	ViewActivationKeys    *bool `json:"view_activation_keys"`
}

// ActivationKeysOp handles communication with the ActivationKey related methods of the
// Red Hat Satellite REST API
type ActivationKeysOp struct {
	client *Client
}

// ActivationKeys is an interface for interacting with
// Red Hat Satellite Activation Keys
type ActivationKeys interface {
	AssociateHostCollections(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error)
	AttachSubscription(ctx context.Context, akID int, subscriptionID int, quantity int) (*ActivationKey, *http.Response, error)
	ContentOverride(ctx context.Context, akID int, contentOverride ActivationKeyContentOverride) (*ActivationKey, *http.Response, error)
	Create(ctx context.Context, akCreate ActivationKeyCreate) (*ActivationKey, *http.Response, error)
	Delete(ctx context.Context, akID int) (*http.Response, error)
	DisassociateHostCollections(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error)
	Get(ctx context.Context, akID int) (*ActivationKey, *http.Response, error)
	List(ctx context.Context, opt *ActivationKeyListOptions) (*ActivationKeyList, *http.Response, error)
	ListByEnvironmentID(ctx context.Context, envID int, opt *ActivationKeyListOptions) (*ActivationKeyList, *http.Response, error)
	ListByOrganizationID(ctx context.Context, orgID int, opt *ActivationKeyListOptions) (*ActivationKeyList, *http.Response, error)
	ListReleases(ctx context.Context, akID int) (*ActivationKeyReleasesList, *http.Response, error)
	Update(ctx context.Context, akID int, akUpdate ActivationKeyUpdate) (*ActivationKey, *http.Response, error)
	UnattachSubscription(ctx context.Context, akID int, subscriptionID int) (*ActivationKey, *http.Response, error)
}

// AssociateHostCollections with an activation key
func (s *ActivationKeysOp) AssociateHostCollections(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/host_collections", activationKeyPath, akID)

	if len(hostCollections) < 1 {
		return nil, nil, NewArgError("hostCollections", "cannot be empty")
	}

	var body struct {
		HostCollectionIDs []int `json:"host_collection_ids"`
	}

	body.HostCollectionIDs = hostCollections

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// AttachSubscription attaches a subscription to an activation key
func (s *ActivationKeysOp) AttachSubscription(ctx context.Context, akID int, subscriptionID int, quantity int) (*ActivationKey, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/add_subscriptions", activationKeyPath, akID)

	var body struct {
		SubscriptionID int `json:"subscription_id"`
		Quantity       int `json:"quantity"`
	}

	body.SubscriptionID = subscriptionID
	body.Quantity = quantity

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// ContentOverride overrides the content an activation key
func (s *ActivationKeysOp) ContentOverride(ctx context.Context, akID int, contentOverride ActivationKeyContentOverride) (*ActivationKey, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/content_override", activationKeyPath, akID)

	if contentOverride.ContentOverrides.Name == nil {
		return nil, nil, NewArgError("contentOverride.Name", "cannot be empty")
	} else if *contentOverride.ContentOverrides.Name == "" {
		return nil, nil, NewArgError("contentOverride.Name", "cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, contentOverride)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// Create a new activation key
func (s *ActivationKeysOp) Create(ctx context.Context, akCreate ActivationKeyCreate) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath

	if akCreate.OrganizationID == nil {
		return nil, nil, NewArgError("akCreate.OrganizationID", "cannot be empty")
	}

	if akCreate.Name == nil {
		return nil, nil, NewArgError("akCreate.Name", "cannot be empty")
	} else if *akCreate.Name == "" {
		return nil, nil, NewArgError("akCreate.Name", "cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, akCreate)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// Delete an activation key by its ID
func (s *ActivationKeysOp) Delete(ctx context.Context, akID int) (*http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DisassociateHostCollections disassociates a list of host collections with an activation key
func (s *ActivationKeysOp) DisassociateHostCollections(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/host_collections", activationKeyPath, akID)

	if len(hostCollections) < 1 {
		return nil, nil, NewArgError("hostCollections", "cannot be empty")
	}

	var body struct {
		HostCollectionIDs []int `json:"host_collection_ids"`
	}

	body.HostCollectionIDs = hostCollections

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// Get a single activation key by its ID
func (s *ActivationKeysOp) Get(ctx context.Context, akID int) (*ActivationKey, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", activationKeyPath, akID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// Performs a list request given a path.
func (s *ActivationKeysOp) list(ctx context.Context, path string) (*ActivationKeyList, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(ActivationKeyList)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, err
}

// List all activation keys or a filtered list of activation keys
func (s *ActivationKeysOp) List(ctx context.Context, opt *ActivationKeyListOptions) (*ActivationKeyList, *http.Response, error) {
	path := activationKeyPath

	if opt.OrganizationID == 0 && opt.EnvironmentID == 0 {
		return nil, nil, NewArgError("Both opt.OrganizationID and opt.EnvironmentID", "cannot be empty")
	}

	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// ListAvailableHostCollections for an activation key
// func (s *ActivationKeysOp) ListAvailableHostCollections(ctx context.Context, akID int, opt ActivationKeyAvailableHostCollectionsListOptions) (*ActivationKeyList, *http.Response, error) {
// 	path := fmt.Sprintf("%s/%d/host_collections/available", activationKeyPath, akID)

// 	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	activationKey := new(ActivationKey)
// 	resp, err := s.client.Do(ctx, req, activationKey)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return activationKey, resp, nil

// }

// ListByEnvironmentID gets all activation keys or a filtered list of activation keys for a specific environment
func (s *ActivationKeysOp) ListByEnvironmentID(ctx context.Context, envID int, opt *ActivationKeyListOptions) (*ActivationKeyList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/activation_keys", katelloEnvironmentsPath, envID)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// ListByOrganizationID gets all activation keys or a filtered list of activation keys for a specific organization
func (s *ActivationKeysOp) ListByOrganizationID(ctx context.Context, orgID int, opt *ActivationKeyListOptions) (*ActivationKeyList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/activation_keys", katelloOrganizationsPath, orgID)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// ListProductContent for an activation key
// func (s *ActivationKeysOp) ListProductContent(ctx context.Context, akID int, opt ActivationKeyListOptions) (*ActivationProductContentList, *http.Response, error) {
// 	path := fmt.Sprintf("%s/%d/releases", activationKeyPath, akID)

// 	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	releases := new(ActivationKeyReleasesList)
// 	resp, err := s.client.Do(ctx, req, releases)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return releases, resp, nil
// }

// ListReleases for an activation key
func (s *ActivationKeysOp) ListReleases(ctx context.Context, akID int) (*ActivationKeyReleasesList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/releases", activationKeyPath, akID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	releases := new(ActivationKeyReleasesList)
	resp, err := s.client.Do(ctx, req, releases)
	if err != nil {
		return nil, resp, err
	}

	return releases, resp, nil
}

// Update an activation key
func (s *ActivationKeysOp) Update(ctx context.Context, akID int, akUpdate ActivationKeyUpdate) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, akUpdate)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// UnattachSubscription detaches a subscription from an activation key
func (s *ActivationKeysOp) UnattachSubscription(ctx context.Context, akID int, subscriptionID int) (*ActivationKey, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/remove_subscriptions", activationKeyPath, akID)

	var body struct {
		SubscriptionID int `json:"subscription_id"`
	}

	body.SubscriptionID = subscriptionID

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}
