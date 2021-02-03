package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const permissionsPath = basePath + "/permissions"

// Permission defines the model of a single permission
type Permission struct {
	ID           *int    `json:"id"`
	Name         *string `json:"name"`
	ResourceType *string `json:"resource_type"`
}

// PermissionsList defines model for a list of permissions.
type PermissionsList struct {
	searchResults
	Results *[]Permission `json:"results"`
}

// PermissionsListOptions specifies the optional parameters to various List methods that
// support pagination.
type PermissionsListOptions struct {
	ListOptions

	// Scope by locations
	LocationID int `url:"location_id,omitempty"`

	// Scope by organizations
	OrganizationID int `url:"organization_id,omitempty"`
}

// PermissionsSearch defines model for searching a list of permissions.
type PermissionsSearch struct {
	Order   *string `json:"order,omitempty"`
	Page    *int    `json:"page,omitempty"`
	PerPage *int    `json:"per_page,omitempty"`
	Search  *string `json:"search,omitempty"`
}

// ResourceTypes defines model for a list of resource types.
type ResourceTypes struct {
	searchResults
	Results *[]struct {
		Name *string `json:"name"`
	} `json:"results"`
}

// Permissions is an interface for interacting with
// Red Hat Satellite permissions
type Permissions interface {
	Get(ctx context.Context, permissionID int) (*Permission, *http.Response, error)
	List(ctx context.Context, opt PermissionsListOptions) (*PermissionsList, *http.Response, error)
}

// PermissionsOp handles communication with the Permissions related methods of the
// Red Hat Satellite REST API
type PermissionsOp struct {
	client *Client
}

// Get a single permission by its ID
func (s *PermissionsOp) Get(ctx context.Context, permissionID int) (*Permission, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", permissionsPath, permissionID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	permission := new(Permission)
	resp, err := s.client.Do(ctx, req, permission)
	if err != nil {
		return nil, resp, err
	}

	return permission, resp, err
}

// List all permissions or a filtered list of permissions
func (s *PermissionsOp) List(ctx context.Context, opt PermissionsListOptions) (*PermissionsList, *http.Response, error) {
	path := permissionsPath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	permissions := new(PermissionsList)
	resp, err := s.client.Do(ctx, req, permissions)
	if err != nil {
		return nil, resp, err
	}

	return permissions, resp, err
}
