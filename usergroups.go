package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const userGroupsPath = basePath + "/usergroups"

// UserGroup defines model for a User Group.
type UserGroup struct {
	Admin              *bool               `json:"admin"`
	CreatedAt          *string             `json:"created_at"`
	UpdatedAt          *string             `json:"updated_at"`
	Name               *string             `json:"name"`
	ID                 *int                `json:"id"`
	ExternalUserGroups *[]genericShortRef  `json:"external_usergroups"`
	UserGroups         *[]genericUserGroup `json:"usergroups"`
	Users              *[]genericUser      `json:"users"`
	Roles              *[]genericRole      `json:"roles"`
}

// UserGroupCreate defines model for the body of the creation of a user group.
type UserGroupCreate struct {
	UserGroup struct {
		Name         *string `json:"name"`
		Admin        *bool   `json:"admin,omitempty"`
		UserIDs      *[]int  `json:"user_ids,omitempty"`
		UserGroupIDs *[]int  `json:"usergroup_ids,omitempty"`
		RoleIDs      *[]int  `json:"role_ids,omitempty"`
	} `json:"usergroup"`
}

// UserGroupUpdate defines model for the body of the update of a user group.
type UserGroupUpdate struct {
	UserGroup struct {
		Name         *string `json:"name,omitempty"`
		Admin        *bool   `json:"admin,omitempty"`
		UserIDs      *[]int  `json:"user_ids,omitempty"`
		UserGroupIDs *[]int  `json:"usergroup_ids,omitempty"`
		RoleIDs      *[]int  `json:"role_ids,omitempty"`
	} `json:"usergroup"`
}

// UserGroups is an interface for interacting with
// Red Hat Satellite roles
type UserGroups interface {
	Create(ctx context.Context, userGroupCreate UserGroupCreate) (*UserGroup, *http.Response, error)
	Delete(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error)
	Get(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error)
	Update(ctx context.Context, userGroupID int, userGroupUpdate UserGroupUpdate) (*UserGroup, *http.Response, error)
}

// UserGroupsOp handles communication with the User Group related methods of the
// Red Hat Satellite REST API
type UserGroupsOp struct {
	client *Client
}

// Create a user group
func (s *UserGroupsOp) Create(ctx context.Context, userGroupCreate UserGroupCreate) (*UserGroup, *http.Response, error) {
	path := rolesPath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, userGroupCreate)
	if err != nil {
		return nil, nil, err
	}

	userGroup := new(UserGroup)
	resp, err := s.client.Do(ctx, req, userGroup)
	if err != nil {
		return nil, resp, err
	}

	return userGroup, resp, err
}

// Delete a user group by its ID
func (s *UserGroupsOp) Delete(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", rolesPath, userGroupID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	userGroup := new(UserGroup)
	resp, err := s.client.Do(ctx, req, userGroup)
	if err != nil {
		return nil, resp, err
	}

	return userGroup, resp, err
}

// Get a single user group by its ID
func (s *UserGroupsOp) Get(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", rolesPath, userGroupID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	userGroup := new(UserGroup)
	resp, err := s.client.Do(ctx, req, userGroup)
	if err != nil {
		return nil, resp, err
	}

	return userGroup, resp, err
}

// Update a user group
func (s *UserGroupsOp) Update(ctx context.Context, userGroupID int, userGroupUpdate UserGroupUpdate) (*UserGroup, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", rolesPath, userGroupID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, userGroupUpdate)
	if err != nil {
		return nil, nil, err
	}

	userGroup := new(UserGroup)
	resp, err := s.client.Do(ctx, req, userGroup)
	if err != nil {
		return nil, resp, err
	}

	return userGroup, resp, err
}
