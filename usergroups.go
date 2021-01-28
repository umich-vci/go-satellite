package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const userGroupsPath = basePath + "/usergroups"

// UserGroup defines model for a User Group.
type UserGroup struct {
	Admin              bool               `json:"admin"`
	CreatedAt          string             `json:"created_at"`
	UpdatedAt          string             `json:"updated_at"`
	Name               string             `json:"name"`
	ID                 int                `json:"id"`
	ExternalUserGroups []genericShortRef  `json:"external_usergroups"`
	UserGroups         []genericUserGroup `json:"usergroups"`
	Users              []genericUser      `json:"users"`
	Roles              []genericRole      `json:"roles"`
}

// UserGroupCreate defines model for the body of the creation of a user group.
type UserGroupCreate struct {
	UserGroup struct {
		Name         string `json:"name"`
		Admin        *bool  `json:"admin,omitempty"`
		UserIDs      *[]int `json:"user_ids,omitempty"`
		UserGroupIDs *[]int `json:"usergroup_ids,omitempty"`
		RoleIDs      *[]int `json:"role_ids,omitempty"`
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
	CreateUserGroup(ctx context.Context, userGroupID int, userGroupCreate UserGroupCreate) (*UserGroup, *http.Response, error)
	DeleteUserGroup(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error)
	GetUserGroupByID(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error)
	UpdateUserGroup(ctx context.Context, userGroupID int, userGroupUpdate UserGroupUpdate) (*UserGroup, *http.Response, error)
}

// UserGroupsOp handles communication with the User Group related methods of the
// Red Hat Satellite REST API
type UserGroupsOp struct {
	client *Client
}

// CreateUserGroup creates a user group
func (s *UserGroupsOp) CreateUserGroup(ctx context.Context, userGroupID int, userGroupCreate UserGroupCreate) (*UserGroup, *http.Response, error) {
	path := rolesPath

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

// DeleteUserGroup deletes a user group by its ID
func (s *UserGroupsOp) DeleteUserGroup(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error) {
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

// GetUserGroupByID gets a single role by its ID
func (s *UserGroupsOp) GetUserGroupByID(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error) {
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

// UpdateUserGroup updates a user group
func (s *UserGroupsOp) UpdateUserGroup(ctx context.Context, userGroupID int, userGroupUpdate UserGroupUpdate) (*UserGroup, *http.Response, error) {
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
