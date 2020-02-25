package gosatellite

import (
	"context"
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

// ProductSearch defines model for searching a list of products.
type ProductSearch struct {
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

// Products is an interface for interacting with
// Red Hat Satellite products
type Products interface {
	ListProductsByOrgID(ctx context.Context, orgID int, prodSearch ProductSearch) (*ProductsList, *http.Response, error)
}

// ProductsOp handles communication with the Product related methods of the
// Red Hat Satellite REST API
type ProductsOp struct {
	client *Client
}

// ListProductsByOrgID gets all products or a filtered list of products for a specific organization
func (s *ProductsOp) ListProductsByOrgID(ctx context.Context, orgID int, prodSearch ProductSearch) (*ProductsList, *http.Response, error) {
	path := productsPath

	prodSearch.OrganizationID = &orgID

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
