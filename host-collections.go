package gosatellite

import (
	"context"
	"net/http"
	"strconv"
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
	CreateHostCollection(ctx context.Context, orgID int, hcCreate HostCollectionCreate) (*HostCollection, *http.Response, error)
	DeleteHostCollection(ctx context.Context, hcID int) (*http.Response, error)
	GetHostCollectionByID(ctx context.Context, hcID int) (*HostCollection, *http.Response, error)
	UpdateHostCollection(ctx context.Context, hcID int, hcUpdate HostCollectionUpdate) (*HostCollection, *http.Response, error)
}

// CreateHostCollection creates a new host collection
func (s *HostCollectionsOp) CreateHostCollection(ctx context.Context, orgID int, hcCreate HostCollectionCreate) (*HostCollection, *http.Response, error) {
	path := organizationsPath + "/host_collections"

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

// DeleteHostCollection deletes a host collection
func (s *HostCollectionsOp) DeleteHostCollection(ctx context.Context, hcID int) (*http.Response, error) {
	path := hostCollectionsPath + "/" + strconv.Itoa(hcID)

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

// GetHostCollectionByID gets a host collection by it's ID
func (s *HostCollectionsOp) GetHostCollectionByID(ctx context.Context, hcID int) (*HostCollection, *http.Response, error) {
	path := hostCollectionsPath + "/" + strconv.Itoa(hcID)

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

// UpdateHostCollection updates a host collection
func (s *HostCollectionsOp) UpdateHostCollection(ctx context.Context, hcID int, hcUpdate HostCollectionUpdate) (*HostCollection, *http.Response, error) {
	path := hostCollectionsPath + "/" + strconv.Itoa(hcID)

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
