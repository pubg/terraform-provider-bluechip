package bluechip_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
)

func NewClusterClient[T bluechip_models.ClusterApiResource[P], P bluechip_models.BaseSpec](client *Client, gvk bluechip_models.GroupVersionKind) *ClusterResourceClient[T, P] {
	return &ClusterResourceClient[T, P]{
		Client: client,
		gvk:    gvk,
	}
}

type ClusterResourceClient[T bluechip_models.ClusterApiResource[P], P bluechip_models.BaseSpec] struct {
	*Client
	gvk bluechip_models.GroupVersionKind
}

func (c *ClusterResourceClient[T, P]) Get(ctx context.Context, name string) (T, error) {
	var data T
	req, err := http.NewRequest("GET", c.JoinUrl(c.gvk.ToApiPath(), name), nil)
	if err != nil {
		return data, err
	}
	_, err = c.DoWithType(req, &data)
	return data, err
}

func (c *ClusterResourceClient[T, P]) Upsert(ctx context.Context, data T) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.JoinUrl(c.gvk.ToApiPath()), bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Do(req)
	body := readBodyForError(resp)
	if err != nil {
		return fmt.Errorf("http request failed: %w, body: %s", err, body)
	}
	if len(body) != 0 {
		tflog.Debug(ctx, "response body: %s", map[string]any{"body": string(body), "resp": resp})
	}
	return nil
}

func (c *ClusterResourceClient[T, P]) Delete(ctx context.Context, name string) error {
	req, err := http.NewRequest("DELETE", c.JoinUrl(c.gvk.ToApiPath(), name), nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	body := readBodyForError(resp)
	if err != nil {
		return fmt.Errorf("http request failed: %w, body: %s", err, body)
	}
	if len(body) != 0 {
		tflog.Debug(ctx, "response body: %s", map[string]any{"body": string(body), "resp": resp})
	}
	return nil
}

func (c *ClusterResourceClient[T, P]) List(ctx context.Context) ([]T, error) {
	var nextToken *string
	var items []T

	for {
		req, err := http.NewRequestWithContext(ctx, "GET", c.JoinUrl(c.gvk.ToApiPath()), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/json")

		var data bluechip_models.ListResponseImpl[T]
		_, err = c.DoWithType(req, &data)
		if err != nil {
			return nil, err
		}

		items = append(items, data.Items...)
		nextToken = data.Metadata.NextToken

		if nextToken == nil {
			break
		}
	}

	return items, nil
}

func (c *ClusterResourceClient[T, P]) Search(ctx context.Context, query []bluechip_models.QueryTerm) ([]T, error) {
	var items []T

	listRequest := &bluechip_models.ListRequest{Items: query}

	for {
		reqBuf, err := json.Marshal(listRequest)
		if err != nil {
			return nil, err
		}
		reqStream := bytes.NewBuffer(reqBuf)

		req, err := http.NewRequestWithContext(ctx, "POST", c.JoinUrl(c.gvk.ToApiPath(), "search"), reqStream)
		if err != nil {
			return nil, err
		}

		if listRequest.NextToken != nil {
			q := req.URL.Query()
			q.Add("nextToken", *listRequest.NextToken)
			req.URL.RawQuery = q.Encode()
		}
		var data bluechip_models.ListResponseImpl[T]
		_, err = c.DoWithType(req, &data)

		items = append(items, data.Items...)
		listRequest.NextToken = data.Metadata.NextToken

		if listRequest.NextToken == nil {
			break
		}
	}
	return items, nil
}
