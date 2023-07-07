package batainer

import (
	"context"
	"fmt"
)

func (c *Client) ListVolumes(ctx context.Context, req *ListVolumesRequest) (res *ListVolumesResponse, err error) {
	if req.ClusterID == "" {
		req.ClusterID = "-"
	}
	if req.AccountID == "" {
		req.AccountID = "-"
	}
	httpRes, err := c.client.R().
		SetResult(&ListVolumesResponse{}).
		Get("/api/v1rc1/clusters/" + req.ClusterID + "/accounts/" + req.AccountID + "/volumes")
	if err != nil {
		return res, fmt.Errorf("http get: %v", err)
	}
	if httpRes.IsError() {
		svcErr, ok := httpRes.Error().(*ServiceError)
		if !ok {
			return nil, fmt.Errorf("http get response: %s", httpRes.Body())
		}
		return nil, fmt.Errorf("http get response: %v", svcErr.Message)
	}
	return httpRes.Result().(*ListVolumesResponse), nil
}

func (c *Client) CreateVolume(ctx context.Context, req *CreateVolumeRequest) (vol *Volume, err error) {
	httpRes, err := c.client.R().
		SetBody(req.Volume).
		SetResult(&Volume{}).
		Post("/api/v1rc1/clusters/" + req.ClusterID + "/accounts/" + req.AccountID + "/volumes")
	if err != nil {
		return nil, fmt.Errorf("http get: %v", err)
	}
	if httpRes.IsError() {
		svcErr, ok := httpRes.Error().(*ServiceError)
		if !ok {
			return nil, fmt.Errorf("http post response: %s", httpRes.Body())
		}
		return nil, fmt.Errorf("http post response: %v", svcErr.Message)
	}
	return httpRes.Result().(*Volume), nil
}

func (c *Client) DeleteVolume(ctx context.Context, req *DeleteVolumeRequest) (err error) {
	httpRes, err := c.client.R().
		Delete("/api/v1rc1/clusters/" + req.ClusterID + "/accounts/" + req.AccountID + "/volumes/" + req.VolumeID)
	if err != nil {
		return fmt.Errorf("http delete: %v", err)
	}
	if httpRes.IsError() {
		svcErr, ok := httpRes.Error().(*ServiceError)
		if !ok {
			return fmt.Errorf("http delete response: %s", httpRes.Body())
		}
		return fmt.Errorf("http delete response: %v", svcErr.Message)
	}
	return nil
}

type ListVolumesRequest struct {
	ClusterID string
	AccountID string
}

type ListVolumesResponse struct {
	Volumes []*Volume `json:"volumes"`
}

type CreateVolumeRequest struct {
	ClusterID string
	AccountID string
	Volume    Volume
}

type DeleteVolumeRequest struct {
	ClusterID string
	AccountID string
	VolumeID  string
}

type Volume struct {
	ID        string         `json:"id"`
	AccountID string         `json:"accountID"`
	ClusterID string         `json:"clusterID"`
	Service   *VolumeService `json:"service,omitempty"`
}

type VolumeService struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  struct {
		WebDAV string `json:"webdav"`
		SFTP   string `json:"sftp"`
	}
}
