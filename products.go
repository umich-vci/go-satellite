package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const productsPath = katelloBasePath + "/products"

type productSyncSummary struct {
	Success *int `json:"success"`
	Warning *int `json:"warning"`
	Pending *int `json:"pending"`
}

type productSyncPlan struct {
	CronExpression *string `json:"cron_expression"`
	Description    *string `json:"description"`
	ID             *int    `json:"id"`
	Interval       *string `json:"interval"`
	Name           *string `json:"name"`
	NextSync       *string `json:"next_sync"`
	SyncDate       *string `json:"sync_date"`
}

// Product defines the model of a single product
type Product struct {
	CpID            *string             `json:"cp_id"`
	Description     *string             `json:"description"`
	GPGKeyID        *int                `json:"gpg_key_id"`
	ID              *int                `json:"id"`
	Label           *string             `json:"label"`
	LastSync        *string             `json:"last_sync"`
	LastSyncWords   *string             `json:"last_sync_words"`
	Name            *string             `json:"name"`
	Organization    *shortOrg           `json:"organization"`
	OrganizationID  *int                `json:"organization_id"`
	ProviderID      *int                `json:"provider_id"`
	RepositoryCount *int                `json:"repository_count"`
	SSLCACertID     *int                `json:"ssl_ca_cert_id"`
	SSLClientCertID *int                `json:"ssl_client_cert_id"`
	SSLClientKeyID  *int                `json:"ssl_client_key_id"`
	SyncPlan        *productSyncPlan    `json:"sync_plan"`
	SyncPlanID      *int                `json:"sync_plan_id"`
	SyncState       *string             `json:"sync_state"`
	SyncSummary     *productSyncSummary `json:"sync_summary"`
}

// ProductsList defines model for a list of products.
type ProductsList struct {
	searchResults
	Error   *string    `json:"error"`
	Results *[]Product `json:"results"`
}

// ProductsListOptions specifies the optional parameters to various List methods that
// support pagination.
type ProductsListOptions struct {
	KatelloListOptions

	// Filter products by organization
	OrganizationID int `url:"organization_id,omitempty"`

	// Filter products by subscription
	SubscriptionID int `url:"subscription_id,omitempty"`

	// Filter products by name
	Name string `url:"name,omitempty"`

	// Return enabled products only
	Enabled bool `url:"enabled,omitempty"`

	// Return custom products only
	Custom bool `url:"custom,omitempty"`

	// Return Red Hat (non-custom) products only
	RedHatOnly bool `url:"redhat_only,omitempty"`

	// Whether to include available content attribute in results
	IncludeAvailableContent bool `url:"include_available_content,omitempty"`

	// Filter products by sync plan id
	SyncPlanID int `url:"sync_plan_id,omitempty"`

	// Interpret specified object to return only Products that can be associated with specified object. Only 'sync_plan' is supported.
	AvailableFor string `url:"available_for,omitempty"`
}

// Products is an interface for interacting with
// Red Hat Satellite products
type Products interface {
	ListByOrgID(ctx context.Context, orgID int, opt *ProductsListOptions) (*ProductsList, *http.Response, error)
	List(ctx context.Context, opt *ProductsListOptions) (*ProductsList, *http.Response, error)
}

// ProductsOp handles communication with the Product related methods of the
// Red Hat Satellite REST API
type ProductsOp struct {
	client *Client
}

// Performs a list request given a path.
func (s *ProductsOp) list(ctx context.Context, path string) (*ProductsList, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(ProductsList)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, err
}

// ListByOrgID all products or a filtered list of products for a specific organization
func (s *ProductsOp) ListByOrgID(ctx context.Context, orgID int, opt *ProductsListOptions) (*ProductsList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/products", katelloOrganizationsPath, orgID)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// List all products or a filtered list of products
func (s *ProductsOp) List(ctx context.Context, opt *ProductsListOptions) (*ProductsList, *http.Response, error) {
	path := productsPath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}
