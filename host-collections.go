package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const hostCollectionsPath = katelloBasePath + "/host_collections"

// HostCollection defines model for a Host Collection.
type HostCollection struct {
	CreatedAt      *string        `json:"created_at"`
	Description    *string        `json:"description"`
	HostIDs        *[]int         `json:"host_ids"`
	ID             *int           `json:"id"`
	MaxHosts       *int           `json:"max_hosts"`
	Name           *string        `json:"name"`
	OrganizationID *int           `json:"organization_id"`
	Permissions    *hcPermissions `json:"permissions"`
	TotalHosts     *int           `json:"total_hosts"`
	UnlimitedHosts *bool          `json:"unlimited_hosts"`
	UpdatedAt      *string        `json:"updated_at"`
}

// HostCollectionCreate defines model for creating a host collection.
type HostCollectionCreate struct {
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
	HostIDs        *[]int  `json:"host_ids,omitempty"`
	MaxHosts       *int    `json:"max_hosts,omitempty"`
	UnlimitedHosts *bool   `json:"unlimited_hosts,omitempty"`
}

// HostCollectionUpdate defines model for updating a host collection.
type HostCollectionUpdate struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	HostIDs        *[]int  `json:"host_ids,omitempty"`
	MaxHosts       *int    `json:"max_hosts,omitempty"`
	UnlimitedHosts *bool   `json:"unlimited_hosts,omitempty"`
}

type hcPermissions struct {
	Deletable *bool `json:"deletable"`
	Editable  *bool `json:"editable"`
}

// HostCollectionsOp handles communication with the Host Collection related methods of the
// Red Hat Satellite REST API
type HostCollectionsOp struct {
	client *Client
}

// HostCollections is an interface for interacting with
// Red Hat Satellite Host Collections
type HostCollections interface {
	Create(ctx context.Context, orgID int, hcCreate HostCollectionCreate) (*HostCollection, *http.Response, error)
	Delete(ctx context.Context, hcID int) (*http.Response, error)
	Get(ctx context.Context, hcID int) (*HostCollection, *http.Response, error)
	Update(ctx context.Context, hcID int, hcUpdate HostCollectionUpdate) (*HostCollection, *http.Response, error)
}

// Create a new host collection
func (s *HostCollectionsOp) Create(ctx context.Context, orgID int, hcCreate HostCollectionCreate) (*HostCollection, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/host_collections", katelloOrganizationsPath, orgID)

	if hcCreate.Name == "" {
		return nil, nil, NewArgError("hcCreate.Name", "cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, hcCreate)
	if err != nil {
		return nil, nil, err
	}
	hc := new(HostCollection)
	resp, err := s.client.Do(ctx, req, hc)
	if err != nil {
		return nil, resp, err
	}

	return hc, resp, err
}

// Delete a host collection
func (s *HostCollectionsOp) Delete(ctx context.Context, hcID int) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d", hostCollectionsPath, hcID)

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

// Get a host collection by its ID
func (s *HostCollectionsOp) Get(ctx context.Context, hcID int) (*HostCollection, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", hostCollectionsPath, hcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	hc := new(HostCollection)
	resp, err := s.client.Do(ctx, req, hc)
	if err != nil {
		return nil, resp, err
	}

	return hc, resp, err
}

// Update a host collection
func (s *HostCollectionsOp) Update(ctx context.Context, hcID int, hcUpdate HostCollectionUpdate) (*HostCollection, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", hostCollectionsPath, hcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, hcUpdate)
	if err != nil {
		return nil, nil, err
	}
	hc := new(HostCollection)
	resp, err := s.client.Do(ctx, req, hc)
	if err != nil {
		return nil, resp, err
	}

	return hc, resp, err
}
