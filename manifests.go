package gosatellite

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// ManifestHistoryItem defines model for a single manifest history
type ManifestHistoryItem struct {
	Created       *string `json:"created"`
	ID            *string `json:"id"`
	Status        *string `json:"status"`
	StatusMessage *string `json:"statusMessage"`
}

// ManifestHistory defines model for the manifest history of an organization
type ManifestHistory struct {
	History *[]ManifestHistoryItem
}

// Manifests is an interface for interacting with
// Red Hat Satellite Subscription Manifests
type Manifests interface {
	DeleteManifest(ctx context.Context, orgID int) (*http.Response, error)
	GetManifestHistory(ctx context.Context, orgID int) (*ManifestHistory, *http.Response, error)
	RefreshManifest(ctx context.Context, orgID int) (*http.Response, error)
	UploadManifest(ctx context.Context, orgID int, repoURL *string, manifest []byte) (*ManifestHistory, *http.Response, error)
}

// ManifestsOp handles communication with the Manifest related methods of the
// Red Hat Satellite REST API
type ManifestsOp struct {
	client *Client
}

// DeleteManifest deletes a manifest for an organization by its ID
func (s *ManifestsOp) DeleteManifest(ctx context.Context, orgID int) (*http.Response, error) {
	path := organizationsPath + "/" + strconv.Itoa(orgID) + "/subscriptions/delete_manifest"

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

// GetManifestHistory gets the manifest history for an organization based on its ID
func (s *ManifestsOp) GetManifestHistory(ctx context.Context, orgID int) (*ManifestHistory, *http.Response, error) {
	path := organizationsPath + "/" + strconv.Itoa(orgID) + "/subscriptions/manifest_history"
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	hist := new(ManifestHistory)
	resp, err := s.client.Do(ctx, req, hist)
	if err != nil {
		return nil, resp, err
	}

	return hist, resp, err
}

// RefreshManifest refreshes the manifest attached to an organization
func (s *ManifestsOp) RefreshManifest(ctx context.Context, orgID int) (*http.Response, error) {
	path := organizationsPath + "/" + strconv.Itoa(orgID) + "/subscriptions/refresh_manifest"

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

// UploadManifest uploads a manifest to an organization
func (s *ManifestsOp) UploadManifest(ctx context.Context, orgID int, repoURL *string, manifest []byte) (*ManifestHistory, *http.Response, error) {
	path := organizationsPath + "/" + strconv.Itoa(orgID) + "/subscriptions/upload"

	var body map[string]io.Reader
	if repoURL != nil {
		body["repository_url"] = strings.NewReader(*repoURL)
	}

	body["content"] = bytes.NewReader(manifest)

	req, err := s.client.NewMultipartFormRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, nil, err
	}

	hist := new(ManifestHistory)
	resp, err := s.client.Do(ctx, req, hist)
	if err != nil {
		return nil, resp, err
	}

	return hist, resp, err
}
