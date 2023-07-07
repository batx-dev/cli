package parauser

import (
	resty "github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func NewClient(baseurl string) *Client {
	c := &Client{
		client: resty.New().SetError(&ServiceError{}).SetBaseURL(baseurl),
	}
	return c
}
