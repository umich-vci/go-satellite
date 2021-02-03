package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const locationsPath = basePath + "/locations"

// Location defines model for a Location.
type Location struct {
	Ancestry              *string             `json:"ancestry"`
	ComputeResources      *[]genericCompRes   `json:"compute_resources"`
	ConfigTemplates       *[]genericTemplate  `json:"config_templates"`
	CreatedAt             *string             `json:"created_at"`
	Description           *string             `json:"description"`
	Domains               *[]genericShortRef  `json:"domains"`
	Environments          *[]genericShortRef  `json:"environments"`
	HostGroups            *[]genericReference `json:"hostgroups"`
	HostsCount            *int                `json:"hosts_count"`
	ID                    *int                `json:"id"`
	Media                 *[]genericShortRef  `json:"media"`
	Name                  *string             `json:"name"`
	Organizations         *[]genericReference `json:"organizations"`
	Parameters            *[]orgParameter     `json:"parameters"`
	ParentID              *int                `json:"parent_id"`
	ParentName            *string             `json:"parent_name"`
	ProvisioningTemplates *[]genericTemplate  `json:"provisioning_templates"`
	Ptables               *[]genericPtables   `json:"ptables"`
	//Realms                *[]genericReference  `json:"realms"`
	SelectAllTypes *[]string            `json:"select_all_types"`
	SmartProxies   *[]genericSmartProxy `json:"smart_proxies"`
	Subnets        *[]genericSubnet     `json:"subnets,omitempty"`
	Title          *string              `json:"title"`
	UpdatedAt      *string              `json:"updated_at"`
	Users          *[]genericUser       `json:"users"`
}

// LocationsList defines model for a list of locations.
type LocationsList struct {
	searchResults
	Results *[]Location `json:"results"`
}

// LocationsListOptions specifies the optional parameters to various List methods that
// support pagination.
type LocationsListOptions struct {
	ListOptions

	// Scope by locations
	LocationID int `url:"location_id,omitempty"`

	// Scope by organizations
	OrganizationID int `url:"organization_id,omitempty"`
}

// LocationCreate defines model for creating a location.
type LocationCreate struct {
	Location struct {
		Name                    *string `json:"name"`
		Description             *string `json:"description,omitempty"`
		UserIDs                 *[]int  `json:"user_ids,omitempty"`
		SmartProxyIDs           *[]int  `json:"smart_proxy_ids,omitempty"`
		ComputeResourceIDs      *[]int  `json:"compute_resource_ids,omitempty"`
		MediumIDs               *[]int  `json:"medium_ids,omitempty"`
		ConfigTemplateIDs       *[]int  `json:"config_template_ids,omitempty"`
		PtableIDs               *[]int  `json:"ptable_ids,omitempty"`
		ProvisioningTemplateIDs *[]int  `json:"provisioning_template_ids,omitempty"`
		DomainIDs               *[]int  `json:"domain_ids,omitempty"`
		RealmIDs                *[]int  `json:"realm_ids,omitempty"`
		HostgroupIDs            *[]int  `json:"hostgroup_ids,omitempty"`
		EnvironmentIDs          *[]int  `json:"environment_ids,omitempty"`
		SubnetIDs               *[]int  `json:"subnet_ids,omitempty"`
		ParentID                *int    `json:"parent_id,omitempty"`
		IgnoreTypes             *[]int  `json:"ignore_types,omitempty"`
	} `json:"location,omitempty"`
}

// LocationUpdate defines model for updating a location.
type LocationUpdate struct {
	Location struct {
		Name                    *string `json:"name,omitempty"`
		Description             *string `json:"description,omitempty"`
		UserIDs                 *[]int  `json:"user_ids,omitempty"`
		SmartProxyIDs           *[]int  `json:"smart_proxy_ids,omitempty"`
		ComputeResourceIDs      *[]int  `json:"compute_resource_ids,omitempty"`
		MediumIDs               *[]int  `json:"medium_ids,omitempty"`
		ConfigTemplateIDs       *[]int  `json:"config_template_ids,omitempty"`
		PtableIDs               *[]int  `json:"ptable_ids,omitempty"`
		ProvisioningTemplateIDs *[]int  `json:"provisioning_template_ids,omitempty"`
		DomainIDs               *[]int  `json:"domain_ids,omitempty"`
		RealmIDs                *[]int  `json:"realm_ids,omitempty"`
		HostgroupIDs            *[]int  `json:"hostgroup_ids,omitempty"`
		EnvironmentIDs          *[]int  `json:"environment_ids,omitempty"`
		SubnetIDs               *[]int  `json:"subnet_ids,omitempty"`
		ParentID                *int    `json:"parent_id,omitempty"`
		IgnoreTypes             *[]int  `json:"ignore_types,omitempty"`
	} `json:"location,omitempty"`
}

// Locations is an interface for interacting with
// Red Hat Satellite Locations
type Locations interface {
	Create(ctx context.Context, locCreate LocationCreate) (*Location, *http.Response, error)
	Delete(ctx context.Context, locationID int) (*http.Response, error)
	Get(ctx context.Context, locationID int) (*Location, *http.Response, error)
	List(ctx context.Context, opt *LocationsListOptions) (*LocationsList, *http.Response, error)
	Update(ctx context.Context, locationID int, update LocationUpdate) (*Location, *http.Response, error)
}

// LocationsOp handles communication with the Locations related methods of the
// Red Hat Satellite REST API
type LocationsOp struct {
	client *Client
}

// Create a new location
func (s *LocationsOp) Create(ctx context.Context, locCreate LocationCreate) (*Location, *http.Response, error) {
	path := locationsPath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, locCreate)
	if err != nil {
		return nil, nil, err
	}
	location := new(Location)
	resp, err := s.client.Do(ctx, req, location)
	if err != nil {
		return nil, resp, err
	}

	return location, resp, err
}

// Delete a location by its ID
func (s *LocationsOp) Delete(ctx context.Context, locationID int) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d", locationsPath, locationID)

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

// Get a single location by its ID
func (s *LocationsOp) Get(ctx context.Context, locationID int) (*Location, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", locationsPath, locationID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	location := new(Location)
	resp, err := s.client.Do(ctx, req, location)
	if err != nil {
		return nil, resp, err
	}

	return location, resp, err
}

// List all locations or a filtered list of locations
func (s *LocationsOp) List(ctx context.Context, opt *LocationsListOptions) (*LocationsList, *http.Response, error) {
	path := locationsPath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	locations := new(LocationsList)
	resp, err := s.client.Do(ctx, req, locations)
	if err != nil {
		return nil, resp, err
	}

	return locations, resp, err
}

// Update the settings of a location by its ID
func (s *LocationsOp) Update(ctx context.Context, locationID int, update LocationUpdate) (*Location, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", locationsPath, locationID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, update)
	if err != nil {
		return nil, nil, err
	}

	location := new(Location)
	resp, err := s.client.Do(ctx, req, location)
	if err != nil {
		return nil, resp, err
	}

	return location, resp, err
}
