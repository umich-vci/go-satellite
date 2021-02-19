package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const contentViewsPath = katelloBasePath + "/content_views"

// ContentViewsListOptions specifies the optional parameters to various List methods that
// support pagination.
type ContentViewsListOptions struct {
	KatelloListOptions

	// organization identifier
	OrganizationID int `url:"organization_id,omitempty"`

	// environment identifier
	EnvironmentID int `url:"environment_id,omitempty"`

	// Filter out default content views
	Nondefault bool `url:"nondefault,omitempty"`

	// Filter out composite content views
	Noncomposite bool `url:"noncomposite,omitempty"`

	// Filter only composite content views
	Composite bool `url:"composite,omitempty"`

	// Do not include this array of content views
	Without []string `url:"without,omitempty"`

	// Name of the content view
	Name string `url:"name,omitempty"`
}

// ContentView defines model for a Content View.
type ContentView struct {
	Composite              *bool              `json:"composite"`
	ComponentIDs           *[]int             `json:"component_ids"`
	Default                *bool              `json:"default"`
	ForcePuppetEnvironment *bool              `json:"force_puppet_environment"`
	VersionCount           *int               `json:"version_count"`
	LatestVersion          *string            `json:"latest_version"`
	AutoPublish            *bool              `json:"auto_publish"`
	SolveDependencies      *bool              `json:"solve_dependencies"`
	RepositoryIDs          *[]int             `json:"repository_ids"`
	ID                     *int               `json:"id"`
	Name                   *string            `json:"name"`
	Label                  *string            `json:"label"`
	Description            *string            `json:"description"`
	OrganizationID         *int               `json:"organization_id"`
	Organization           *shortOrg          `json:"organization"`
	CreatedAt              *string            `json:"created_at"`
	UpdatedAt              *string            `json:"updated_at"`
	Environments           *[]shortLE         `json:"environments"`
	Repositories           *[]shortRepository `json:"repositories"`
	Versions               *[]cvVersions      `json:"versions"`
	ActivationKeys         *[]genericShortRef `json:"activation_keys"`
	NextVersion            *string            `json:"next_version"`
	LastPublished          *string            `json:"last_published"`
	Permissions            *cvPermissions     `json:"permissions"`
	//PuppetModules          *[]unknown         `json:"puppet_modules"`
	//Components             *[]unknown         `json:"components"`
	//ContentViewComponents  *[]unknown         `json:"content_view_components"`
}

type cvVersions struct {
	ID             *int    `json:"id"`
	Version        *string `json:"version"`
	Published      *string `json:"published"`
	EnvironmentIDs *[]int  `json:"environment_ids"`
}

type cvPermissions struct {
	ViewContentViews            *bool `json:"view_content_views"`
	EditContentViews            *bool `json:"edit_content_views"`
	DestroyContentViews         *bool `json:"destroy_content_views"`
	PublishContentViews         *bool `json:"publish_content_views"`
	PromoteOrRemoveContentViews *bool `json:"promote_or_remove_content_views"`
}

// ContentViewsList defines model for a list of Content Views.
type ContentViewsList struct {
	searchResults
	Results *[]ContentView `json:"results"`
}

// ContentViews is an interface for interacting with
// Red Hat Satellite Content Views
type ContentViews interface {
	List(ctx context.Context, opt *ContentViewsListOptions) (*ContentViewsList, *http.Response, error)
	ListByOrganizationID(ctx context.Context, orgID int, opt *ContentViewsListOptions) (*ContentViewsList, *http.Response, error)
}

// ContentViewsOp handles communication with the Content Views related methods of the
// Red Hat Satellite REST API
type ContentViewsOp struct {
	client *Client
}

// Performs a list request given a path.
func (s *ContentViewsOp) list(ctx context.Context, path string) (*ContentViewsList, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(ContentViewsList)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, err
}

// List all Content Views or a filtered list of Content Views
func (s *ContentViewsOp) List(ctx context.Context, opt *ContentViewsListOptions) (*ContentViewsList, *http.Response, error) {
	path := contentViewsPath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// ListByOrganizationID all Content Views in an organization or a filtered list of Content Views in an organization
func (s *ContentViewsOp) ListByOrganizationID(ctx context.Context, orgID int, opt *ContentViewsListOptions) (*ContentViewsList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/content_views", katelloOrganizationsPath, orgID)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}
