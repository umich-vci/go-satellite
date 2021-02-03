package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

// ManifestHistoryItem defines model for a single manifest history
type ManifestHistoryItem struct {
	Created       *string `json:"created"`
	ID            *string `json:"id"`
	Status        *string `json:"status"`
	StatusMessage *string `json:"statusMessage"`
}

// ManifestUpload defines model for the response from a manifest upload to an organization
type ManifestUpload struct {
	ID        *string `json:"id"`
	Label     *string `json:"label"`
	Pending   *bool   `json:"pending"`
	Action    *string `json:"action"`
	UserName  *string `json:"username"`
	StartedAt *string `json:"started_at"`
	EndedAt   *string `json:"ended_at"`
	State     *string `json:"state"`
	Result    *string `json:"result"`
}

// Manifests is an interface for interacting with
// Red Hat Satellite Subscription Manifests
type Manifests interface {
	Delete(ctx context.Context, orgID int) (*http.Response, error)
	GetHistory(ctx context.Context, orgID int) (*[]ManifestHistoryItem, *http.Response, error)
	Refresh(ctx context.Context, orgID int) (*http.Response, error)
	Upload(ctx context.Context, orgID int, repoURL *string, manifest []byte, manifestFilename string) (*ManifestUpload, *http.Response, error)
}

// ManifestsOp handles communication with the Manifest related methods of the
// Red Hat Satellite REST API
type ManifestsOp struct {
	client *Client
}

// Delete a manifest for an organization by its ID
func (s *ManifestsOp) Delete(ctx context.Context, orgID int) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d/subscriptions/delete_manifest", katelloOrganizationsPath, orgID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// GetHistory of a manifest for an organization based on its ID
func (s *ManifestsOp) GetHistory(ctx context.Context, orgID int) (*[]ManifestHistoryItem, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/subscriptions/manifest_history", katelloOrganizationsPath, orgID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	hist := new([]ManifestHistoryItem)
	resp, err := s.client.Do(ctx, req, hist)
	if err != nil {
		return nil, resp, err
	}

	return hist, resp, err
}

// Refresh the manifest attached to an organization
func (s *ManifestsOp) Refresh(ctx context.Context, orgID int) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d/subscriptions/refresh_manifest", katelloOrganizationsPath, orgID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Upload a manifest to an organization
func (s *ManifestsOp) Upload(ctx context.Context, orgID int, repoURL *string, manifest []byte, manifestFilename string) (*ManifestUpload, *http.Response, error) {
	path := fmt.Sprintf("%s/%d/subscriptions/upload", katelloOrganizationsPath, orgID)

	req, err := s.client.NewManifestUploadRequest(ctx, http.MethodPost, path, manifest, manifestFilename)
	if err != nil {
		return nil, nil, err
	}

	status := new(ManifestUpload)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, err
}
