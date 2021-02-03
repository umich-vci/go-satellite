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

// ExternalUserGroup defines model for an External User Group.
type ExternalUserGroup struct {
	ID             *int                   `json:"id"`
	Name           *string                `json:"name"`
	AuthSourceLDAP *genericAuthSourceLDAP `json:"auth_source_ldap"`
}

// ExternalUserGroup2 defines model for an External User Group.
type ExternalUserGroup2 struct {
	ID           *int    `json:"id"`
	Name         *string `json:"name"`
	AuthSourceID *int    `json:"auth_source_id"`
	UserGroupID  *int    `json:"usergroup_id"`
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

// ExternalUserGroupCreate defines model for the body of the creation of an external user group.
type ExternalUserGroupCreate struct {
	ExternalUserGroup struct {
		Name         string `json:"name"`
		AuthSourceID int    `json:"auth_source_id"`
	} `json:"external_usergroup"`
}

// ExternalUserGroupUpdate defines model for the body of the update of an external user group.
type ExternalUserGroupUpdate struct {
	ExternalUserGroup struct {
		Name         *string `json:"name,omitempty"`
		AuthSourceID *int    `json:"auth_source_id,omitempty"`
	} `json:"external_usergroup"`
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
	CreateExternalUserGroup(ctx context.Context, userGroupID int, externalUserGroupCreate ExternalUserGroupCreate) (*ExternalUserGroup, *http.Response, error)
	CreateUserGroup(ctx context.Context, userGroupCreate UserGroupCreate) (*UserGroup, *http.Response, error)
	DeleteExternalUserGroup(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup2, *http.Response, error)
	DeleteUserGroup(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error)
	GetExternalUserGroupByID(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup, *http.Response, error)
	GetUserGroupByID(ctx context.Context, userGroupID int) (*UserGroup, *http.Response, error)
	UpdateExternalUserGroup(ctx context.Context, userGroupID int, externalUserGroupID int, externalUserGroupUpdate ExternalUserGroupUpdate) (*ExternalUserGroup, *http.Response, error)
	UpdateUserGroup(ctx context.Context, userGroupID int, userGroupUpdate UserGroupUpdate) (*UserGroup, *http.Response, error)
}

// UserGroupsOp handles communication with the User Group related methods of the
// Red Hat Satellite REST API
type UserGroupsOp struct {
	client *Client
}

// CreateExternalUserGroup creates an external user group linked to a user group
func (s *UserGroupsOp) CreateExternalUserGroup(ctx context.Context, userGroupID int, externalUserGroupCreate ExternalUserGroupCreate) (*ExternalUserGroup, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/external_usergroups", rolesPath, userGroupID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, externalUserGroupCreate)
	if err != nil {
		return nil, nil, err
	}

	externalUserGroup := new(ExternalUserGroup)
	resp, err := s.client.Do(ctx, req, externalUserGroup)
	if err != nil {
		return nil, resp, err
	}

	return externalUserGroup, resp, err
}

// CreateUserGroup creates a user group
func (s *UserGroupsOp) CreateUserGroup(ctx context.Context, userGroupCreate UserGroupCreate) (*UserGroup, *http.Response, error) {
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

// DeleteExternalUserGroup deletes an external user group by its ID
func (s *UserGroupsOp) DeleteExternalUserGroup(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup2, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/external_usergroups/%d", rolesPath, userGroupID, externalUserGroupID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	externalUserGroup := new(ExternalUserGroup2)
	resp, err := s.client.Do(ctx, req, externalUserGroup)
	if err != nil {
		return nil, resp, err
	}

	return externalUserGroup, resp, err
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

// GetExternalUserGroupByID gets a external user group by its ID
func (s *UserGroupsOp) GetExternalUserGroupByID(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/external_usergroups/%d", rolesPath, userGroupID, externalUserGroupID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	externalUserGroup := new(ExternalUserGroup)
	resp, err := s.client.Do(ctx, req, externalUserGroup)
	if err != nil {
		return nil, resp, err
	}

	return externalUserGroup, resp, err
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

// UpdateExternalUserGroup updates a user group
func (s *UserGroupsOp) UpdateExternalUserGroup(ctx context.Context, userGroupID int, externalUserGroupID int, externalUserGroupUpdate ExternalUserGroupUpdate) (*ExternalUserGroup, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/external_usergroups/%d", rolesPath, userGroupID, externalUserGroupID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, externalUserGroupUpdate)
	if err != nil {
		return nil, nil, err
	}

	externalUserGroup := new(ExternalUserGroup)
	resp, err := s.client.Do(ctx, req, externalUserGroup)
	if err != nil {
		return nil, resp, err
	}

	return externalUserGroup, resp, err
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
