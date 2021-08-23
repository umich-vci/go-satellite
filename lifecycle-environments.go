package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const katelloEnvironmentsPath = katelloBasePath + "/environments"
const environmentsPath = basePath + "/environments"

// LifecycleEnvironmentsListOptions specifies the optional parameters to various List methods that
// support pagination.
type LifecycleEnvironmentsListOptions struct {
	KatelloListOptions

	// organization identifier
	OrganizationID int `url:"organization_id,omitempty"`

	// set true if you want to see only library environments
	Library bool `url:"library,omitempty"`

	// filter only environments containing this name
	Name string `url:"name,omitempty"`
}

// LifecycleEnvironment defines model for a Lifecycle Environment.
type LifecycleEnvironment struct {
	Library                     *bool                       `json:"library"`
	RegistryNamePattern         *string                     `json:"registry_name_pattern"`
	RegistryUnauthenticatedPull *bool                       `json:"registry_unauthenticated_pull"`
	ID                          *int                        `json:"id"`
	Name                        *string                     `json:"name"`
	Label                       *string                     `json:"label"`
	Description                 *string                     `json:"description"`
	OrganizationID              *int                        `json:"organization_id"`
	Organization                *shortOrg                   `json:"organization"`
	CreatedAt                   *string                     `json:"created_at"`
	UpdatedAt                   *string                     `json:"updated_at"`
	Prior                       *genericShortRef            `json:"prior"`
	Successor                   *genericShortRef            `json:"successor"`
	Counts                      *lifecycleEnvironmentCounts `json:"counts"`
	Permissions                 *lePermissions              `json:"permissions"`
}

type lifecycleEnvironmentCounts struct {
	ContentHosts       *int            `json:"content_hosts"`
	ContentViews       *int            `json:"content_views"`
	Packages           *int            `json:"packages"`
	PuppetModules      *int            `json:"puppet_modules"`
	ModuleStreams      *int            `json:"module_streams"`
	Errata             *leErrataCounts `json:"errata"`
	YumRepositories    *int            `json:"yum_repositories"`
	DockerRepositories *int            `json:"docker_repositories"`
	OSTreeRepositories *int            `json:"ostree_repositories"`
	Products           *int            `json:"products"`
}

type leErrataCounts struct {
	Security    *int `json:"security"`
	Bugfix      *int `json:"bugfix"`
	Enhancement *int `json:"enhancement"`
	Total       *int `json:"total"`
}

type lePermissions struct {
	CreateLifecycleEnvironments               *bool `json:"create_lifecycle_environments"`
	ViewLifecycleEnvironments                 *bool `json:"view_lifecycle_environments"`
	EditLifecycleEnvironments                 *bool `json:"edit_lifecycle_environments"`
	DestroyLifecycleEnvironments              *bool `json:"destroy_lifecycle_environments"`
	PromoteOrRemoveContentViewsToEnvironments *bool `json:"promote_or_remove_content_views_to_environments"`
}

// LifecycleEnvironmentsList defines model for a list of Lifecycle Environments.
type LifecycleEnvironmentsList struct {
	searchResults
	Results *[]LifecycleEnvironment `json:"results"`
}

// LifecycleEnvironments is an interface for interacting with
// Red Hat Satellite Lifecycle Environments
type LifecycleEnvironments interface {
	List(ctx context.Context, opt *LifecycleEnvironmentsListOptions) (*LifecycleEnvironmentsList, *http.Response, error)
	ListByOrganizationID(ctx context.Context, orgID int, opt *LifecycleEnvironmentsListOptions) (*LifecycleEnvironmentsList, *http.Response, error)
}

// LifecycleEnvironmentsOp handles communication with the Lifecycle Environments related methods of the
// Red Hat Satellite REST API
type LifecycleEnvironmentsOp struct {
	client *Client
}

// Performs a list request given a path.
func (s *LifecycleEnvironmentsOp) list(ctx context.Context, path string) (*LifecycleEnvironmentsList, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(LifecycleEnvironmentsList)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, err
}

// List all Lifecycle Environments or a filtered list of Lifecycle Environments
func (s *LifecycleEnvironmentsOp) List(ctx context.Context, opt *LifecycleEnvironmentsListOptions) (*LifecycleEnvironmentsList, *http.Response, error) {
	path := katelloEnvironmentsPath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// ListByOrganizationID all Lifecycle Environments in an organization or a filtered list of Lifecycle Environments in an organization
func (s *LifecycleEnvironmentsOp) ListByOrganizationID(ctx context.Context, orgID int, opt *LifecycleEnvironmentsListOptions) (*LifecycleEnvironmentsList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/environments", katelloOrganizationsPath, orgID)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}
