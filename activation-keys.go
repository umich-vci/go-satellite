package gosatellite

import (
	"context"
	"net/http"
	"strconv"
)

const activationKeyPath = katelloBasePath + "/activation_keys"

// ActivationKey defines model for an Activation Key.
type ActivationKey struct {
	AutoAttach       *bool                 `json:"auto_attach"`
	ContentOverrides *[]akContentOverrides `json:"content_overrides"`
	ContentView      *genericShortRef      `json:"content_view"`
	ContentViewID    *int                  `json:"content_view_id"`
	CreatedAt        *string               `json:"created_at"`
	Description      *string               `json:"description"`
	Environment      *genericShortRef      `json:"environment"`
	EnvironmentID    *int                  `json:"environment_id"`
	HostCollections  *[]genericShortRef    `json:"host_collections"`
	ID               *int                  `json:"id"`
	MaxHosts         *int                  `json:"max_hosts"`
	Name             *string               `json:"name"`
	Organization     *shortOrg             `json:"organization"`
	OrganizationID   *int                  `json:"organization_id"`
	Permissions      *akPermissions        `json:"permissions"`
	Products         *[]genericShortRef    `json:"products"`
	ReleaseVersion   *string               `json:"release_version"`
	ServiceLevel     *string               `json:"service_level"`
	UnlimitedHosts   *bool                 `json:"unlimited_hosts"`
	UpdatedAt        *string               `json:"updated_at"`
	UsageCount       *int                  `json:"usage_count"`
	UserID           *int                  `json:"user_id"`
}

// ActivationKeyList defines model for a list of activation keys.
type ActivationKeyList struct {
	searchResults
	Error   *string          `json:"error"`
	Results *[]ActivationKey `json:"results"`
}

// ActivationKeySearch defines model for searching a list of activation keys.
type ActivationKeySearch struct {
	OrganizationID *int    `json:"organization_id,omitempty"`
	EnvironmentID  *int    `json:"environment_id,omitempty"`
	ContentViewID  *int    `json:"content_view_id,omitempty"`
	Name           *string `json:"name,omitempty"`
	Search         *string `json:"search,omitempty"`
	Page           *int    `json:"page,omitempty"`
	PerPage        *int    `json:"per_page,omitempty"`
	Order          *string `json:"order,omitempty"`
	FullResult     *bool   `json:"full_result,omitempty"`
	SortBy         *string `json:"sort_by,omitempty"`
	SortOrder      *string `json:"sort_order,omitempty"`
}

// ActivationKeyCreate defines model for creating an activation key.
type ActivationKeyCreate struct {
	OrganizationID *int    `json:"organization_id"`
	Name           *string `json:"name"`
	Description    *string `json:"description,omitempty"`
	EnvironmentID  *int    `json:"environment_id,omitempty"`
	ContentViewID  *int    `json:"content_view_id,omitempty"`
	MaxHosts       *int    `json:"max_hosts,omitempty"`
	UnlimitedHosts *bool   `json:"unlimited_hosts,omitempty"`
}

// ActivationKeyUpdate defines model for updating an activation key.
type ActivationKeyUpdate struct {
	OrganizationID *int    `json:"organization_id,omitempty"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	EnvironmentID  *int    `json:"environment_id,omitempty"`
	ContentViewID  *int    `json:"content_view_id,omitempty"`
	MaxHosts       *int    `json:"max_hosts,omitempty"`
	UnlimitedHosts *bool   `json:"unlimited_hosts,omitempty"`
	ReleaseVersion *string `json:"release_version,omitempty"`
	ServiceLevel   *string `json:"service_level,omitempty"`
	AutoAttach     *bool   `json:"auto_attach,omitempty"`
}

// ActivationKeyContentOverride defines model for defining a content override for an activation key.
type ActivationKeyContentOverride struct {
	ContentOverrides struct {
		ContentLabel *string `json:"content_label"`
		Value        *string `json:"value,omitempty"`
		Name         *string `json:"name,omitempty"`
		Remove       *bool   `json:"remove,omitempty"`
	} `json:"content_overrides"`
}

type akContentOverrides struct {
	ContentLabel *string `json:"content_label"`
	Name         *string `json:"name"`
	Value        *string `json:"value"`
}

type akPermissions struct {
	DestroyActivationKeys *bool `json:"destroy_activation_keys"`
	EditActivationKeys    *bool `json:"edit_activation_keys"`
	ViewActivationKeys    *bool `json:"view_activation_keys"`
}

// ActivationKeysOp handles communication with the ActivationKey related methods of the
// Red Hat Satellite REST API
type ActivationKeysOp struct {
	client *Client
}

// ActivationKeys is an interface for interacting with
// Red Hat Satellite Activation Keys
type ActivationKeys interface {
	AssociateHostCollectionsWithActivationKey(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error)
	CreateActivationKey(ctx context.Context, akCreate ActivationKeyCreate) (*ActivationKey, *http.Response, error)
	DeleteActivationKey(ctx context.Context, akID int) (*http.Response, error)
	DisassociateHostCollectionsWithActivationKey(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error)
	GetActivationKeyByID(ctx context.Context, akID int) (*ActivationKey, *http.Response, error)
	UpdateActivationKey(ctx context.Context, akID int, akUpdate ActivationKeyUpdate) (*ActivationKey, *http.Response, error)
}

// AttachSubscriptionToActivationKey attaches a subscription to an activation key
func (s *ActivationKeysOp) AttachSubscriptionToActivationKey(ctx context.Context, akID int, subscriptionID int, quantity int) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID) + "/add_subscriptions"

	var body struct {
		SubscriptionID int `json:"subscription_id"`
		Quantity       int `json:"quantity"`
	}

	body.SubscriptionID = subscriptionID
	body.Quantity = quantity

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// AssociateHostCollectionsWithActivationKey associates a list of host collections with an activation key
func (s *ActivationKeysOp) AssociateHostCollectionsWithActivationKey(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID) + "/host_collections"

	if len(hostCollections) < 1 {
		return nil, nil, NewArgError("hostCollections", "cannot be empty")
	}

	var body struct {
		HostCollectionIDs []int `json:"host_collection_ids"`
	}

	body.HostCollectionIDs = hostCollections

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// UnattachSubscriptionFromActivationKey detaches a subscription from an activation key
func (s *ActivationKeysOp) UnattachSubscriptionFromActivationKey(ctx context.Context, akID int, subscriptionID int) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID) + "/remove_subscriptions"

	var body struct {
		SubscriptionID int `json:"subscription_id"`
	}

	body.SubscriptionID = subscriptionID

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// DisassociateHostCollectionsWithActivationKey disassociates a list of host collections with an activation key
func (s *ActivationKeysOp) DisassociateHostCollectionsWithActivationKey(ctx context.Context, akID int, hostCollections []int) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID) + "/host_collections"

	if len(hostCollections) < 1 {
		return nil, nil, NewArgError("hostCollections", "cannot be empty")
	}

	var body struct {
		HostCollectionIDs []int `json:"host_collection_ids"`
	}

	body.HostCollectionIDs = hostCollections

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// CreateActivationKey creates a new activation key
func (s *ActivationKeysOp) CreateActivationKey(ctx context.Context, akCreate ActivationKeyCreate) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath

	if akCreate.OrganizationID == nil {
		return nil, nil, NewArgError("akCreate.OrganizationID", "cannot be empty")
	}

	if akCreate.Name == nil {
		return nil, nil, NewArgError("akCreate.Name", "cannot be empty")
	} else if *akCreate.Name == "" {
		return nil, nil, NewArgError("akCreate.Name", "cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, akCreate)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// DeleteActivationKey deletes an activation key by its ID
func (s *ActivationKeysOp) DeleteActivationKey(ctx context.Context, akID int) (*http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID)

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

// GetActivationKeyByID gets a single activation key by its ID
func (s *ActivationKeysOp) GetActivationKeyByID(ctx context.Context, akID int) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// ListActivationKeysByOrgID gets all activation keys or a filtered list of activation keys for a specific organization
func (s *ActivationKeysOp) ListActivationKeysByOrgID(ctx context.Context, orgID int, akSearch ActivationKeySearch) (*ActivationKeyList, *http.Response, error) {
	path := organizationsPath + strconv.Itoa(orgID) + "/activation_keys"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, akSearch)
	if err != nil {
		return nil, nil, err
	}

	activationKeys := new(ActivationKeyList)
	resp, err := s.client.Do(ctx, req, activationKeys)
	if err != nil {
		return nil, resp, err
	}

	return activationKeys, resp, nil
}

// ListActivationKeys gets all activation keys or a filtered list of activation keys
func (s *ActivationKeysOp) ListActivationKeys(ctx context.Context, akSearch ActivationKeySearch) (*ActivationKeyList, *http.Response, error) {
	path := activationKeyPath

	if akSearch.OrganizationID == nil && akSearch.EnvironmentID == nil {
		return nil, nil, NewArgError("Both akSearch.OrganizationID and akSearch.EnvironmentID", "cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, akSearch)
	if err != nil {
		return nil, nil, err
	}

	activationKeys := new(ActivationKeyList)
	resp, err := s.client.Do(ctx, req, activationKeys)
	if err != nil {
		return nil, resp, err
	}

	return activationKeys, resp, nil
}

// UpdateActivationKey updates an activation key
func (s *ActivationKeysOp) UpdateActivationKey(ctx context.Context, akID int, akUpdate ActivationKeyUpdate) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, akUpdate)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}

// OverrideContentForActivationKey overrides the content an activation key
func (s *ActivationKeysOp) OverrideContentForActivationKey(ctx context.Context, akID int, contentOverride ActivationKeyContentOverride) (*ActivationKey, *http.Response, error) {
	path := activationKeyPath + "/" + strconv.Itoa(akID) + "/content_override"

	if contentOverride.ContentOverrides.Name == nil {
		return nil, nil, NewArgError("contentOverride.Name", "cannot be empty")
	} else if *contentOverride.ContentOverrides.Name == "" {
		return nil, nil, NewArgError("contentOverride.Name", "cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, contentOverride)
	if err != nil {
		return nil, nil, err
	}

	activationKey := new(ActivationKey)
	resp, err := s.client.Do(ctx, req, activationKey)
	if err != nil {
		return nil, resp, err
	}

	return activationKey, resp, nil
}
