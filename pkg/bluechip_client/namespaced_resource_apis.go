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
	req, err := http.NewRequest("GET", c.JoinUrl(c.gvk.ToApiPath(), namespace, name), nil)
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
	body := readBodyForError(resp)
	if err != nil {
		return fmt.Errorf("http request failed: %w, body: %s", err, body)
	}
	if len(body) != 0 {
		tflog.Debug(ctx, "response body: %s", map[string]any{"body": string(body), "resp": resp})
	}
	return nil
}

func (c *NamespacedResourceClient[T, P]) Delete(ctx context.Context, namespace string, name string) error {
	req, err := http.NewRequest("DELETE", c.JoinUrl(c.gvk.ToApiPath(), namespace, name), nil)
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

func (c *NamespacedResourceClient[T, P]) List(ctx context.Context, namespace string) ([]T, error) {
	panic("implement me")
}

func (c *NamespacedResourceClient[T, P]) Search(ctx context.Context, namespace string, query any) {
	panic("implement me")
}
