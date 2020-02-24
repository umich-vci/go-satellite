package gosatellite

import (
	"context"
	"net/http"
	"strconv"
	"strings"
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
	GetPermissionByID(ctx context.Context, permissionID int) (*Permission, *http.Response, error)
	ListPermissions(ctx context.Context, name *string, resourceType *string, page *int, perPage *int) (*PermissionsList, *http.Response, error)
}

// PermissionsOp handles communication with the Permissions related methods of the
// Red Hat Satellite REST API
type PermissionsOp struct {
	client *Client
}

// GetPermissionByID gets a single permission by its ID
func (s *PermissionsOp) GetPermissionByID(ctx context.Context, permissionID int) (*Permission, *http.Response, error) {
	path := permissionsPath + "/" + strconv.Itoa(permissionID)

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

// ListPermissions gets all permissions or a filtered list of permissions
func (s *PermissionsOp) ListPermissions(ctx context.Context, name *string, resourceType *string, page *int, perPage *int) (*PermissionsList, *http.Response, error) {
	path := permissionsPath

	queryStrings := []string{}

	if name != nil {
		if *name != "" {
			queryStrings = append(queryStrings, "name="+*name)
		}
	}

	if resourceType != nil {
		if *resourceType != "" {
			queryStrings = append(queryStrings, "resource_type="+*resourceType)
		}
	}

	if page != nil {
		if *page < 0 {
			return nil, nil, NewArgError("page", "must be greater than 0")
		}
		queryStrings = append(queryStrings, "page="+strconv.Itoa(*page))

	}

	if perPage != nil {
		if *perPage < 0 {
			return nil, nil, NewArgError("perPage", "must be greater than 0")
		}
		queryStrings = append(queryStrings, "per_page="+strconv.Itoa(*perPage))
	}

	if len(queryStrings) > 0 {
		path = path + "?" + strings.Join(queryStrings, "&")
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
