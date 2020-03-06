package gosatellite

import (
	"context"
	"net/http"
	"strconv"
)

const organizationsPath = katelloBasePath + "/organizations"

type orgOwnerDetails struct {
	AutobindDisabled      *bool                     `json:"autobindDisabled"`
	ContentAccessMode     *string                   `json:"contentAccessMode"`
	ContentAccessModeList *string                   `json:"contentAccessModeList"`
	ContentPrefix         *string                   `json:"contentPrefix"`
	Created               *string                   `json:"created"`
	DefaultServiceLevel   *string                   `json:"defaultServiceLevel"`
	DisplayName           *string                   `json:"displayName"`
	Href                  *string                   `json:"href"`
	ID                    *string                   `json:"id"`
	Key                   *string                   `json:"key"`
	LastRefreshed         *string                   `json:"lastRefreshed"`
	LogLevel              *string                   `json:"logLevel"`
	ParentOwner           *string                   `json:"parentOwner"`
	Updated               *string                   `json:"updated"`
	UpstreamConsumer      *orgOwnerUpstreamConsumer `json:"upstreamConsumer"`
	VirtWho               *bool                     `json:"virt_who"`
}

type orgOwnerUpstreamConsumer struct {
	APIURL            *string                         `json:"apiUrl"`
	ContentAccessMode *string                         `json:"contentAccessMode"`
	Created           *string                         `json:"created"`
	ID                *string                         `json:"id"`
	IDCert            *orgOwnerUpstreamConsumerIDCert `json:"idCert"`
	Name              *string                         `json:"name"`
	OwnerID           *string                         `json:"ownerId"`
	Type              *orgOwnerUpstreamConsumerType   `json:"type"`
	Updated           *string                         `json:"updated"`
	UUID              *string                         `json:"uuid"`
	WebURL            *string                         `json:"webUrl"`
}

type orgOwnerUpstreamConsumerIDCert struct {
	Cert    *string                               `json:"cert"`
	Created *string                               `json:"created"`
	ID      *string                               `json:"id"`
	Key     *string                               `json:"key"`
	Serial  *orgOwnerUpstreamConsumerIDCertSerial `json:"serial"`
	Updated *string                               `json:"updated"`
}

type orgOwnerUpstreamConsumerIDCertSerial struct {
	Collected  *bool   `json:"collected"`
	Created    *string `json:"created"`
	Expiration *string `json:"expiration"`
	ID         *int    `json:"id"`
	Revoked    *bool   `json:"revoked"`
	Serial     *int    `json:"serial"`
	Updated    *string `json:"updated"`
}

type orgOwnerUpstreamConsumerType struct {
	ID       *string `json:"id"`
	Label    *string `json:"label"`
	Manifest *bool   `json:"manifest"`
}

type orgParameter struct {
	CreatedAt     *string `json:"created_at"`
	ID            *int    `json:"id"`
	Name          *string `json:"name"`
	ParameterType *string `json:"parameter_type"`
	Priority      *int    `json:"priority"`
	UpdatedAt     *string `json:"updated_at"`
	Value         *string `json:"value"`
}

// Organization defines model for an Organization.
type Organization struct {
	Ancestry              *string             `json:"ancestry"`
	ComputeResources      *[]genericCompRes   `json:"compute_resources"`
	ConfigTemplates       *[]genericTemplate  `json:"config_templates"`
	CreatedAt             *string             `json:"created_at"`
	DefaultContentViewID  *int                `json:"default_content_view_id"`
	Description           *string             `json:"description"`
	Domains               *[]genericShortRef  `json:"domains"`
	Environments          *[]genericShortRef  `json:"environments"`
	HostGroups            *[]genericReference `json:"hostgroups"`
	HostsCount            *int                `json:"hosts_count"`
	ID                    *int                `json:"id"`
	Label                 *string             `json:"label"`
	LibraryID             *int                `json:"library_id"`
	Locations             *[]genericReference `json:"locations"`
	Media                 *[]genericShortRef  `json:"media"`
	Name                  *string             `json:"name"`
	OwnerDetails          *orgOwnerDetails    `json:"owner_details"`
	Parameters            *[]orgParameter     `json:"parameters"`
	ParentID              *int                `json:"parent_id"`
	ParentName            *string             `json:"parent_name"`
	ProvisioningTemplates *[]genericTemplate  `json:"provisioning_templates"`
	Ptables               *[]genericPtables   `json:"ptables"`
	//Realms                *[]genericReference  `json:"realms"`
	RedHatRepositoryURL *string              `json:"redhat_repository_url"`
	SelectAllTypes      *[]string            `json:"select_all_types"`
	ServiceLevel        *string              `json:"service_level"`
	ServiceLevels       *[]string            `json:"service_levels"`
	SmartProxies        *[]genericSmartProxy `json:"smart_proxies"`
	Subnets             *[]genericSubnet     `json:"subnets,omitempty"`
	//SystemPurposes
	Title     *string        `json:"title"`
	UpdatedAt *string        `json:"updated_at"`
	Users     *[]genericUser `json:"users"`
}

// OrganizationCreate defines model for OrganizationCreate.
type OrganizationCreate struct {
	Organization struct {
		ComputeResourceIDs      *[]int  `json:"compute_resource_ids,omitempty"`
		ConfigTemplateIDs       *[]int  `json:"config_template_ids,omitempty"`
		Description             *string `json:"description,omitempty"`
		DomainIDs               *[]int  `json:"domain_ids,omitempty"`
		EnvironmentIDs          *[]int  `json:"environment_ids,omitempty"`
		HostgroupIDs            *[]int  `json:"hostgroup_ids,omitempty"`
		Label                   *string `json:"label,omitempty"`
		MediumIDs               *[]int  `json:"medium_ids,omitempty"`
		Name                    string  `json:"name"`
		ProvisioningTemplateIDs *[]int  `json:"provisioning_template_ids,omitempty"`
		PtableIDs               *[]int  `json:"ptable_ids,omitempty"`
		RealmIds                *[]int  `json:"realm_ids,omitempty"`
		SmartProxyIDs           *[]int  `json:"smart_proxy_ids,omitempty"`
		SubnetIDs               *[]int  `json:"subnet_ids,omitempty"`
		UserIDs                 *[]int  `json:"user_ids,omitempty"`
	} `json:"organization"`
}

// OrganizationUpdate defines model for OrganizationUpdate.
type OrganizationUpdate struct {
	Organization struct {
		ComputeResourceIDs      *[]int    `json:"compute_resource_ids,omitempty"`
		ConfigTemplateIDs       *[]int    `json:"config_template_ids,omitempty"`
		Description             *string   `json:"description,omitempty"`
		DomainIDs               *[]int    `json:"domain_ids,omitempty"`
		EnvironmentIDs          *[]int    `json:"environment_ids,omitempty"`
		HostgroupIDs            *[]int    `json:"hostgroup_ids,omitempty"`
		IgnoreTypes             *[]string `json:"ignore_types,omitempty"`
		MediumIds               *[]int    `json:"medium_ids,omitempty"`
		Name                    *string   `json:"name,omitempty"`
		ParentID                *int      `json:"parent_id,omitempty"`
		ProvisioningTemplateIDs *[]int    `json:"provisioning_template_ids,omitempty"`
		PtableIDs               *[]int    `json:"ptable_ids,omitempty"`
		RealmIDs                *[]int    `json:"realm_ids,omitempty"`
		SmartProxyIDs           *[]int    `json:"smart_proxy_ids,omitempty"`
		SubnetIDs               *[]int    `json:"subnet_ids,omitempty"`
		UserIDs                 *[]int    `json:"user_ids,omitempty"`
	} `json:"organization"`
	RedhatRepositoryURL *string `json:"redhat_repository_url,omitempty"`
}

// Organizations is an interface for interacting with
// Red Hat Satellite organizations
type Organizations interface {
	CreateOrganization(ctx context.Context, orgCreate OrganizationCreate) (*Organization, *http.Response, error)
	DeleteOrganization(ctx context.Context, orgID int) (*http.Response, error)
	GetOrganizationByID(ctx context.Context, orgID int) (*Organization, *http.Response, error)
	UpdateOrganization(ctx context.Context, orgID int, update OrganizationUpdate) (*Organization, *http.Response, error)
}

// OrganizationsOp handles communication with the Organization related methods of the
// Red Hat Satellite REST API
type OrganizationsOp struct {
	client *Client
}

// CreateOrganization creates a new organization
func (s *OrganizationsOp) CreateOrganization(ctx context.Context, orgCreate OrganizationCreate) (*Organization, *http.Response, error) {
	path := organizationsPath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, orgCreate)
	if err != nil {
		return nil, nil, err
	}
	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}

	return org, resp, err
}

// DeleteOrganization deletes an organization by its ID
func (s *OrganizationsOp) DeleteOrganization(ctx context.Context, orgID int) (*http.Response, error) {
	path := organizationsPath + "/" + strconv.Itoa(orgID)

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

// GetOrganizationByID gets a single Organization by its ID
func (s *OrganizationsOp) GetOrganizationByID(ctx context.Context, orgID int) (*Organization, *http.Response, error) {
	path := organizationsPath + "/" + strconv.Itoa(orgID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}

	return org, resp, err
}

// UpdateOrganization the settings of an organization by its ID
func (s *OrganizationsOp) UpdateOrganization(ctx context.Context, orgID int, update OrganizationUpdate) (*Organization, *http.Response, error) {
	path := organizationsPath + "/" + strconv.Itoa(orgID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, update)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}

	return org, resp, err
}
