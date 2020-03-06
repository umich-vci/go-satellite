package gosatellite

import (
	"context"
	"net/http"
	"strconv"
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

// LocationsSearch defines model for searching a list of locations.
type LocationsSearch struct {
	Search  *string `json:"search,omitempty"`
	Order   *string `json:"order,omitempty"`
	Page    *int    `json:"page,omitempty"`
	PerPage *int    `json:"per_page,omitempty"`
}

// Locations is an interface for interacting with
// Red Hat Satellite Locations
type Locations interface {
	CreateLocation(ctx context.Context, locCreate LocationCreate) (*Location, *http.Response, error)
	DeleteLocation(ctx context.Context, locationID int) (*http.Response, error)
	GetLocationByID(ctx context.Context, locationID int) (*Location, *http.Response, error)
	ListLocations(ctx context.Context, locSearch LocationsSearch) (*LocationsList, *http.Response, error)
	UpdateLocation(ctx context.Context, locationID int, update LocationUpdate) (*Location, *http.Response, error)
}

// LocationsOp handles communication with the Locations related methods of the
// Red Hat Satellite REST API
type LocationsOp struct {
	client *Client
}

// CreateLocation creates a new location
func (s *LocationsOp) CreateLocation(ctx context.Context, locCreate LocationCreate) (*Location, *http.Response, error) {
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

// DeleteLocation deletes a location by its ID
func (s *LocationsOp) DeleteLocation(ctx context.Context, locationID int) (*http.Response, error) {
	path := locationsPath + "/" + strconv.Itoa(locationID)

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

// GetLocationByID gets a single location by its ID
func (s *LocationsOp) GetLocationByID(ctx context.Context, locationID int) (*Location, *http.Response, error) {
	path := locationsPath + "/" + strconv.Itoa(locationID)

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

// ListLocations gets all locations or a filtered list of locations
func (s *LocationsOp) ListLocations(ctx context.Context, locSearch LocationsSearch) (*LocationsList, *http.Response, error) {
	path := locationsPath

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, locSearch)
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

// UpdateLocation the settings of a location by its ID
func (s *LocationsOp) UpdateLocation(ctx context.Context, locationID int, update LocationUpdate) (*Location, *http.Response, error) {
	path := locationsPath + "/" + strconv.Itoa(locationID)

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
