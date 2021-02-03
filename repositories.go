package gosatellite

import (
	"context"
	"fmt"
	"net/http"
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

// RepositoriesListOptions specifies the optional parameters to various List methods that
// support pagination.
type RepositoriesListOptions struct {
	KatelloListOptions

	// ID of an organization to show repositories in
	OrganizationID int `url:"organization_id,omitempty"`

	// ID of a product to show repositories of
	ProductID int `url:"product_id,omitempty"`

	// ID of an environment to show repositories in
	EnvironmentID int `url:"environment_id,omitempty"`

	// ID of a content view to show repositories in
	ContentViewID int `url:"content_view_id,omitempty"`

	// ID of a content view version to show repositories in
	ContentViewVersionID int `url:"content_view_version_id,omitempty"`

	// Id of a deb package to find repositories that contain the deb
	DebID int `url:"deb_id,omitempty"`

	// Id of an erratum to find repositories that contain the erratum
	ErratumID int `url:"erratum_id,omitempty"`

	// Id of a rpm package to find repositories that contain the rpm
	RpmID int `url:"rpm_id,omitempty"`

	// Id of a file to find repositories that contain the file
	FileID int `url:"file_id,omitempty"`

	// Id of an ansible collection to find repositories that contain the ansible collection
	AnsibleCollectionID int `url:"ansible_collection_id,omitempty"`

	// Id of an ostree branch to find repositories that contain that branch
	OSTreeBranchID int `url:"ostree_branch_id,omitempty"`

	// show repositories in Library and the default content view
	Library bool `url:"library,omitempty"`

	// show archived repositories
	Archived bool `url:"archived,omitempty"`

	// limit to only repositories of this type
	// Must be one of: puppet, deb, ansible_collection, ostree, docker, yum, file.
	ContentType string `url:"content_type,omitempty"`

	// name of the repository
	Name string `url:"name,omitempty"`

	// label of the repository
	Label string `url:"label,omitempty"`

	// description of the repository
	Description string `url:"description,omitempty"`

	// interpret specified object to return only Repositories that can be associated with specified object.
	// Only 'content_view' & 'content_view_version' are supported.
	AvailableFor string `url:"available_for,omitempty"`

	// only repositories having at least one of the specified content type ex: rpm , erratum
	// Must be one of: puppet_module, deb, ansible collection, ostree, docker_manifest, docker_manifest_list, docker_tag, docker_blob, rpm, modulemd, erratum, distribution, package_category, package_group, yum_repo_metadata_file, srpm, file.
	WithContent string `url:"with_content,omitempty"`
}

// Repositories is an interface for interacting with
// Red Hat Satellite repositories
type Repositories interface {
	Get(ctx context.Context, repoID int) (*Repository, *http.Response, error)
	List(ctx context.Context, opt RepositoriesListOptions) (*RepositoriesList, *http.Response, error)
}

// RepositoriesOp handles communication with the Repository related methods of the
// Red Hat Satellite REST API
type RepositoriesOp struct {
	client *Client
}

// Get a single repository by its ID
func (s *RepositoriesOp) Get(ctx context.Context, repoID int) (*Repository, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", repositoriesPath, repoID)

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

// List all repositories or a filtered list of repositories
func (s *RepositoriesOp) List(ctx context.Context, opt RepositoriesListOptions) (*RepositoriesList, *http.Response, error) {
	path := repositoriesPath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	repositories := new(RepositoriesList)
	resp, err := s.client.Do(ctx, req, repositories)
	if err != nil {
		return nil, resp, err
	}

	return repositories, resp, err
}
