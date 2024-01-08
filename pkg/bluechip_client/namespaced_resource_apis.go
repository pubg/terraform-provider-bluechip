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

func NewNamespacedClient[T bluechip_models.NamespacedApiResource[P], P bluechip_models.BaseSpec](client *Client, gvk bluechip_models.GroupVersionKind) *NamespacedResourceClient[T, P] {
	return &NamespacedResourceClient[T, P]{
		Client: client,
		gvk:    gvk,
	}
}

type NamespacedResourceClient[T bluechip_models.NamespacedApiResource[P], P bluechip_models.BaseSpec] struct {
	*Client
	gvk bluechip_models.GroupVersionKind
}

func (c *NamespacedResourceClient[T, P]) Get(ctx context.Context, namespace string, name string) (T, error) {
	var data T
	req, err := http.NewRequestWithContext(ctx, "GET", c.JoinUrl(c.gvk.ToApiPath(), namespace, name), nil)
	if err != nil {
		return data, err
	}
	_, err = c.DoWithType(req, &data)
	return data, err
}

func (c *NamespacedResourceClient[T, P]) Upsert(ctx context.Context, namespace string, object T) error {
	reqBuf, err := json.Marshal(object)
	if err != nil {
		return err
	}

	reqStream := bytes.NewBuffer(reqBuf)
	req, err := http.NewRequestWithContext(ctx, "PUT", c.JoinUrl(c.gvk.ToApiPath(), namespace), reqStream)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Do(req)
	body := ReadBodyForError(resp)
	if err != nil {
		return fmt.Errorf("http request failed: %w, body: %s", err, body)
	}
	if len(body) != 0 {
		tflog.Debug(ctx, "response body: %s", map[string]any{"body": string(body), "resp": resp})
	}
	return nil
}

func (c *NamespacedResourceClient[T, P]) Delete(ctx context.Context, namespace string, name string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", c.JoinUrl(c.gvk.ToApiPath(), namespace, name), nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	body := ReadBodyForError(resp)
	if err != nil {
		return fmt.Errorf("http request failed: %w, body: %s", err, body)
	}
	if len(body) != 0 {
		tflog.Debug(ctx, "response body: %s", map[string]any{"body": string(body), "resp": resp})
	}
	return nil
}

func (c *NamespacedResourceClient[T, P]) List(ctx context.Context, namespace string) ([]T, error) {
	var nextToken *string
	var items []T

	for {
		req, err := http.NewRequestWithContext(ctx, "GET", c.JoinUrl(c.gvk.ToApiPath(), namespace), nil)
		if err != nil {
			return nil, err
		}

		if nextToken != nil {
			q := req.URL.Query()
			q.Add("nextToken", *nextToken)
			req.URL.RawQuery = q.Encode()
		}
		var data bluechip_models.ListResponseImpl[T]
		_, err = c.DoWithType(req, &data)

		items = append(items, data.Items...)
		nextToken = data.Metadata.NextToken

		if nextToken == nil {
			break
		}
	}

	return items, nil
}

func (c *NamespacedResourceClient[T, P]) Search(ctx context.Context, namespace string, query []bluechip_models.QueryTerm) ([]T, error) {
	var items []T

	listRequest := &bluechip_models.ListRequest{Items: query}

	for {
		reqBuf, err := json.Marshal(listRequest)
		if err != nil {
			return nil, err
		}
		reqStream := bytes.NewBuffer(reqBuf)

		req, err := http.NewRequestWithContext(ctx, "POST", c.JoinUrl(c.gvk.ToApiPath(), namespace, "search"), reqStream)
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
		listRequest.NextToken = data.Metadata.NextToken

		if listRequest.NextToken == nil {
			break
		}
	}
	return items, nil
}
