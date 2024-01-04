package bluechip_client

import (
	"context"
	"net/http"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
)

func (c *Client) Whoami(ctx context.Context) (*bluechip_models.WhoamiResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.JoinUrl("/auth/whoami"), nil)
	if err != nil {
		return nil, err
	}

	var resp bluechip_models.WhoamiResponse
	if _, err := c.DoWithType(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) ApiResources(ctx context.Context) (*bluechip_models.ApiResourcesResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.JoinUrl("/apis"), nil)
	if err != nil {
		return nil, err
	}

	var resp bluechip_models.ApiResourcesResponse
	if _, err := c.DoWithType(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
