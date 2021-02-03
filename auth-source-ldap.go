package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const authSourceLDAPsPath = basePath + "/auth_source_ldaps"

// AuthSourceLDAP defines model for an LDAP authentication source.
type AuthSourceLDAP struct {
	Host               *string             `json:"host"`
	Port               *int                `json:"port"`
	Account            *string             `json:"account"`
	BaseDN             *string             `json:"base_dn"`
	LDAPFilter         *string             `json:"ldap_filter"`
	AttrLogin          *string             `json:"attr_login"`
	AttrFirstName      *string             `json:"attr_firstname"`
	AttrLastName       *string             `json:"attr_lastname"`
	AttrMail           *string             `json:"attr_mail"`
	AttrPhoto          *string             `json:"attr_photo"`
	OnTheFlyRegister   *bool               `json:"onthefly_register"`
	UserGroupSync      *bool               `json:"usergroup_sync"`
	TLS                *bool               `json:"tls"`
	ServerType         *string             `json:"server_type"`
	GroupsBase         *string             `json:"groups_base"`
	UseNetGroups       *bool               `json:"use_netgroups"`
	CreatedAt          *string             `json:"created_at"`
	UpdatedAt          *string             `json:"updated_at"`
	ID                 *int                `json:"id"`
	Type               *string             `json:"type"`
	Name               *string             `json:"name"`
	ExternalUserGroups *[]genericShortRef  `json:"external_usergroups"`
	Locations          *[]genericReference `json:"locations"`
	Organizations      *[]genericReference `json:"organizations"`
}

// AuthSourceLDAPList defines model for a list of LDAP authentication sources.
type AuthSourceLDAPList struct {
	searchResults
	Results *[]AuthSourceLDAP `json:"results"`
}

// AuthSourceLDAPsListOptions specifies the optional parameters to various List methods that
// support pagination.
type AuthSourceLDAPsListOptions struct {
	ListOptions

	// Scope by locations
	LocationID int `url:"location_id,omitempty"`

	// Scope by organizations
	OrganizationID int `url:"organization_id,omitempty"`
}

// AuthSourceLDAPs is an interface for interacting with
// Red Hat Satellite Auth Source LDAPs
type AuthSourceLDAPs interface {
	List(ctx context.Context, opt *AuthSourceLDAPsListOptions) (*AuthSourceLDAPList, *http.Response, error)
	ListByLocationID(ctx context.Context, locID int, opt *AuthSourceLDAPsListOptions) (*AuthSourceLDAPList, *http.Response, error)
	ListByOrganizationID(ctx context.Context, orgID int, opt *AuthSourceLDAPsListOptions) (*AuthSourceLDAPList, *http.Response, error)
}

// AuthSourceLDAPsOp handles communication with the LDAP authentication source related methods of the
// Red Hat Satellite REST API
type AuthSourceLDAPsOp struct {
	client *Client
}

// Performs a list request given a path.
func (s *AuthSourceLDAPsOp) list(ctx context.Context, path string) (*AuthSourceLDAPList, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(AuthSourceLDAPList)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, err
}

// List all LDAP Authentication Sources or a filtered list of LDAP Authentication Sources
func (s *AuthSourceLDAPsOp) List(ctx context.Context, opt *AuthSourceLDAPsListOptions) (*AuthSourceLDAPList, *http.Response, error) {
	path := authSourceLDAPsPath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// ListByLocationID all LDAP Authentication Sources or a filtered list of LDAP Authentication Sources
func (s *AuthSourceLDAPsOp) ListByLocationID(ctx context.Context, locID int, opt *AuthSourceLDAPsListOptions) (*AuthSourceLDAPList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/auth_source_ldaps", locationsPath, locID)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}

// ListByOrganizationID all LDAP Authentication Sources or a filtered list of LDAP Authentication Sources
func (s *AuthSourceLDAPsOp) ListByOrganizationID(ctx context.Context, orgID int, opt *AuthSourceLDAPsListOptions) (*AuthSourceLDAPList, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/auth_source_ldaps", organizationsPath, orgID)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.list(ctx, path)
}
