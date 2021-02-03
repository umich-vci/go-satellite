package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

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

// ExternalUserGroups is an interface for interacting with
// Red Hat Satellite external user groups
type ExternalUserGroups interface {
	Create(ctx context.Context, userGroupID int, externalUserGroupCreate ExternalUserGroupCreate) (*ExternalUserGroup, *http.Response, error)
	Delete(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup2, *http.Response, error)
	Get(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup, *http.Response, error)
	Update(ctx context.Context, userGroupID int, externalUserGroupID int, externalUserGroupUpdate ExternalUserGroupUpdate) (*ExternalUserGroup, *http.Response, error)
}

// ExternalUserGroupsOp handles communication with the External User Group related methods of the
// Red Hat Satellite REST API
type ExternalUserGroupsOp struct {
	client *Client
}

// Create an external user group linked to a user group
func (s *ExternalUserGroupsOp) Create(ctx context.Context, userGroupID int, externalUserGroupCreate ExternalUserGroupCreate) (*ExternalUserGroup, *http.Response, error) {
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

// Delete an external user group by its ID
func (s *ExternalUserGroupsOp) Delete(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup2, *http.Response, error) {
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

// Get an external user group by its ID
func (s *ExternalUserGroupsOp) Get(ctx context.Context, userGroupID int, externalUserGroupID int) (*ExternalUserGroup, *http.Response, error) {
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

// Update an external user group
func (s *ExternalUserGroupsOp) Update(ctx context.Context, userGroupID int, externalUserGroupID int, externalUserGroupUpdate ExternalUserGroupUpdate) (*ExternalUserGroup, *http.Response, error) {
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
