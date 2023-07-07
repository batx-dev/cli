package batainer

import (
	"context"
	"fmt"
)

func (c *Client) ListInstanceTypes(ctx context.Context, req *ListInstanceTypesRequest) (res *ListInstanceTypesResponse, err error) {
	if req == nil {
		req = &ListInstanceTypesRequest{}
	}
	if req.ClusterID == "" {
		req.ClusterID = "-"
	}
	httpRes, err := c.client.R().
		SetResult(&ListInstanceTypesResponse{}).
		Get("/api/v1rc1/clusters/" + req.ClusterID + "/instanceTypes")
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
	return httpRes.Result().(*ListInstanceTypesResponse), nil
}

type ListInstanceTypesRequest struct {
	ClusterID string
}

type ListInstanceTypesResponse struct {
	InstanceTypes []*InstanceType `json:"instanceTypes"`
}

type InstanceType struct {
	ID        string `json:"id"`
	ClusterID string `json:"clusterID"`
	CPU       struct {
		Cores uint   `json:"cores"`
		Model string `json:"model"`
	}
	Memory struct {
		SizeGB uint `json:"sizeGB"`
	}
	GPU struct {
		Count uint   `json:"count"`
		Model string `json:"model"`
	}
}
