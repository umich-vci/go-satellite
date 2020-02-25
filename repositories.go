package gosatellite

import (
	"context"
	"net/http"
	"strconv"
)

const repositoriesPath = katelloBasePath + "/repositories"

type repoContentView struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type repoContentCounts struct {
	Deb                *int `json:"deb"`
	DockerManifest     *int `json:"docker_manifest"`
	DockerManifestList *int `json:"docker_manifest_list"`
	DockerTag          *int `json:"docker_tag"`
	Erratum            *int `json:"erratum"`
	File               *int `json:"file"`
	ModuleStream       *int `json:"module_stream"`
	OSTreeBranch       *int `json:"ostree_branch"`
	Package            *int `json:"package"`
	PackageGroup       *int `json:"package_group"`
	PuppetModule       *int `json:"puppet_module"`
	RPM                *int `json:"rpm"`
	SRPM               *int `json:"srpm"`
}

type repoEnvironment struct {
	ID                          *int  `json:"id"`
	RegistryUnauthenticatedPull *bool `json:"registry_unauthenticated_pull"`
}

type repoLastSync struct {
	EndedAt   *string  `json:"ended_at"`
	ID        *int     `json:"id"`
	Progress  *float64 `json:"progress"`
	Result    *string  `json:"result"`
	StartedAt *string  `json:"started_at"`
	State     *string  `json:"state"`
	Username  *string  `json:"username"`
}

type repoPermissions struct {
	Deletable *bool `json:"deletable"`
}

type repoProduct struct {
	CpID     *string              `json:"cp_id"`
	ID       *int                 `json:"id"`
	Name     *string              `json:"name"`
	Orphaned *bool                `json:"orphaned"`
	RedHat   *bool                `json:"redhat"`
	SyncPlan *repoProductSyncPlan `json:"sync_plan"`
}

type repoProductSyncPlan struct {
	Description *string `json:"description"`
	Interval    *string `json:"interval"`
	Name        *string `json:"name"`
	NextSync    *string `json:"next_sync"`
	SyncDate    *string `json:"sync_date"`
}

// Repository defines the model of a single repository
type Repository struct {
	Arch                            *string            `json:"arch"`
	BackendIdentifier               *string            `json:"backend_identifier"`
	ChecksumType                    *string            `json:"checksum_type"`
	ComputedOSTreeUpstreamSyncDepth *int               `json:"computed_ostree_upstream_sync_depth"`
	ContainerRepositoryName         *string            `json:"container_repository_name"`
	ContentCounts                   *repoContentCounts `json:"content_counts"`
	ContentID                       *string            `json:"content_id"`
	ContentLabel                    *string            `json:"content_label"`
	ContentType                     *string            `json:"content_type"`
	ContentView                     *repoContentView   `json:"content_view"`
	ContentViewVersionID            *string            `json:"content_view_version_id"`
	CreatedAt                       *string            `json:"created_at"`
	//"deb_architectures": null,
	//"deb_components": null,
	//"deb_releases": null,
	Description *string `json:"description"`
	// "docker_tags_whitelist": null,
	// "docker_upstream_name": null,
	DownloadPolicy *string          `json:"download_policy"`
	Environment    *repoEnvironment `json:"environment"`
	FullPath       *string          `json:"full_path"`
	//"gpg_key": null,
	//"gpg_key_id": null,
	ID *int `json:"id"`
	//"ignorable_content": null,
	IgnoreGlobalProxy *bool         `json:"ignore_global_proxy"`
	Label             *string       `json:"label"`
	LastSync          *repoLastSync `json:"last_sync"`
	LastSyncWords     *string       `json:"last_sync_words"`
	//"library_instance_id": null,
	Major        *int      `json:"major"`
	Minor        *string   `json:"minor"`
	MirrorOnSync *bool     `json:"mirror_on_sync"`
	Name         *int      `json:"name"`
	Organization *shortOrg `json:"organization"`
	//"ostree_branches": [],
	//"ostree_upstream_sync_depth": null,
	//"ostree_upstream_sync_policy": null,
	Permissions  *repoPermissions `json:"permissions"`
	Product      *repoProduct     `json:"product"`
	Promoted     *bool            `json:"promoted"`
	RelativePath *string          `json:"relative_path"`
	// "ssl_ca_cert": {
	//     "id": null,
	//     "name": null
	// },
	// "ssl_ca_cert_id": null,
	// "ssl_client_cert": {
	//     "id": null,
	//     "name": null
	// },
	// "ssl_client_cert_id": null,
	// "ssl_client_key": {
	//     "id": null,
	//     "name": null
	// },
	// "ssl_client_key_id": null,
	Unprotected            *bool   `json:"unprotected"`
	UpdatedAt              *string `json:"updated_at"`
	UpstreamAuthExists     *bool   `json:"upstream_auth_exists"`
	UpstreamPasswordExists *bool   `json:"upstream_password_exists"`
	//"upstream_username": null,
	URL             *string `json:"url"`
	VerifySSLOnSync *bool   `json:"verify_ssl_on_sync"`
}

// RepositoriesList defines model for a list of repositories.
type RepositoriesList struct {
	searchResults
	Error   *string       `json:"error"`
	Results *[]Repository `json:"results"`
}

// RepositorySearch defines model for searching a list of repositories.
type RepositorySearch struct {
	OrganizationID          *int    `json:"organization_id,omitempty"`
	SubscriptionID          *int    `json:"subscription_id,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Enabled                 *bool   `json:"enabled,omitempty"`
	Custom                  *bool   `json:"custom,omitempty"`
	RedHatOnly              *bool   `json:"redhat_only,omitempty"`
	IncludeAvailableContent *bool   `json:"include_available_content,omitempty"`
	SyncPlanID              *int    `json:"sync_plan_id,omitempty"`
	AvailableFor            *string `json:"available_for,omitempty"`
	Search                  *string `json:"search,omitempty"`
	Page                    *int    `json:"page,omitempty"`
	PerPage                 *int    `json:"per_page,omitempty"`
	Order                   *string `json:"order,omitempty"`
	FullResult              *bool   `json:"full_result,omitempty"`
	SortBy                  *string `json:"sort_by,omitempty"`
	SortOrder               *string `json:"sort_order,omitempty"`
}

// Repositories is an interface for interacting with
// Red Hat Satellite repositories
type Repositories interface {
	GetRepositoryByID(ctx context.Context, repoID int) (*Repository, *http.Response, error)
	ListRepositories(ctx context.Context, prodSearch ProductSearch) (*ProductsList, *http.Response, error)
}

// RepositoriesOp handles communication with the Repository related methods of the
// Red Hat Satellite REST API
type RepositoriesOp struct {
	client *Client
}

// GetRepositoryByID gets a single repository by its ID
func (s *RepositoriesOp) GetRepositoryByID(ctx context.Context, repoID int) (*Repository, *http.Response, error) {
	path := repositoriesPath + "/" + strconv.Itoa(repoID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	repo := new(Repository)
	resp, err := s.client.Do(ctx, req, repo)
	if err != nil {
		return nil, resp, err
	}

	return repo, resp, err
}

// ListRepositories gets all repositories or a filtered list of repositories
func (s *RepositoriesOp) ListRepositories(ctx context.Context, prodSearch ProductSearch) (*ProductsList, *http.Response, error) {
	path := productsPath

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, prodSearch)
	if err != nil {
		return nil, nil, err
	}

	products := new(ProductsList)
	resp, err := s.client.Do(ctx, req, products)
	if err != nil {
		return nil, resp, err
	}

	return products, resp, err
}
