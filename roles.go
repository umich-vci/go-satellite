package gosatellite

import (
	"context"
	"net/http"
	"strconv"
)

const rolesPath = basePath + "/roles"

// Role defines model for a Role.
type Role struct {
	Builtin       *int                  `json:"builtin"`
	ClonedFromID  *int                  `json:"cloned_from_id"`
	Description   *string               `json:"description"`
	Filters       *[]genericIDReference `json:"filters"`
	ID            *int                  `json:"id"`
	Locations     *[]genericReference   `json:"locations"`
	Name          *string               `json:"name"`
	Organizations *[]genericReference   `json:"organizations"`
	Origin        *string               `json:"origin"`
}

// RoleCreate defines model for the body of the creation of a role.
type RoleCreate struct {
	Role struct {
		Name            *string `json:"name"`
		Description     *string `json:"description,omitempty"`
		LocationIDs     *[]int  `json:"location_ids,omitempty"`
		OrganizationIDs *[]int  `json:"organization_ids,omitempty"`
	} `json:"role"`
}

// RoleUpdate defines model for the body of the update of a role.
type RoleUpdate struct {
	Role struct {
		Name            *string `json:"name,omitempty"`
		Description     *string `json:"description,omitempty"`
		LocationIDs     *[]int  `json:"location_ids,omitempty"`
		OrganizationIDs *[]int  `json:"organization_ids,omitempty"`
	} `json:"role"`
}

// Roles is an interface for interacting with
// Red Hat Satellite roles
type Roles interface {
	CreateRole(ctx context.Context, roleCreate RoleCreate) (*Role, *http.Response, error)
	DeleteRole(ctx context.Context, roleID int) (*http.Response, error)
	GetRoleByID(ctx context.Context, roleID int) (*Role, *http.Response, error)
	UpdateRole(ctx context.Context, roleID int, roleUpdate RoleUpdate) (*Role, *http.Response, error)
}

// RolesOp handles communication with the Role related methods of the
// Red Hat Satellite REST API
type RolesOp struct {
	client *Client
}

// CreateRole creates a new role
func (s *RolesOp) CreateRole(ctx context.Context, roleCreate RoleCreate) (*Role, *http.Response, error) {
	path := rolesPath

	if roleCreate.Role.Name == nil {
		return nil, nil, NewArgError("roleCreate.Role.Name", "cannot be empty")
	}

	if *roleCreate.Role.Name == "" {
		return nil, nil, NewArgError("roleCreate.Role.Name", "cannot be an empty string")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, roleCreate)
	if err != nil {
		return nil, nil, err
	}
	role := new(Role)
	resp, err := s.client.Do(ctx, req, role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, err
}

// DeleteRole deletes a single role by its ID
func (s *RolesOp) DeleteRole(ctx context.Context, roleID int) (*http.Response, error) {
	path := rolesPath + "/" + strconv.Itoa(roleID)
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

// GetRoleByID gets a single role by its ID
func (s *RolesOp) GetRoleByID(ctx context.Context, roleID int) (*Role, *http.Response, error) {
	path := rolesPath + "/" + strconv.Itoa(roleID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	role := new(Role)
	resp, err := s.client.Do(ctx, req, role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, err
}

// UpdateRole updates a role
func (s *RolesOp) UpdateRole(ctx context.Context, roleID int, roleUpdate RoleUpdate) (*Role, *http.Response, error) {
	path := rolesPath + "/" + strconv.Itoa(roleID)

	if roleUpdate.Role.Name != nil {
		if *roleUpdate.Role.Name == "" {
			return nil, nil, NewArgError("roleUpdate.Role.Name", "cannot be an empty string")
		}
	}

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, roleUpdate)
	if err != nil {
		return nil, nil, err
	}
	role := new(Role)
	resp, err := s.client.Do(ctx, req, role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, err
}
