package bluechip_authenticator

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwlog"
)

type Client struct {
	client  *http.Client
	version string
	address string
}

func NewClient(client *http.Client, version string, address string) *Client {
	return &Client{client: client, version: version, address: address}
}

func (c *Client) doLogin(req *http.Request) (*LoginResponse, *http.Response, error) {
	req.Header.Set("User-Agent", "bluechip-go-http/"+c.version)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode/100 != 2 {
		bodyBuf := bluechip_client.ReadBodyForError(resp)
		tflog.Debug(context.Background(), "Login failed", fwlog.Field("status_code", resp.StatusCode), fwlog.Field("body", string(bodyBuf)), fwlog.Field("request", req))
		return nil, resp, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBuf))
	}

	var loginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return nil, resp, err
	}

	return &loginResponse, resp, nil
}

func (c *Client) LoginWithBasic(ctx context.Context, username string, password string) (string, error) {
	u, err := url.JoinPath(c.address, "/auth/users/login")
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", u, nil)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
	lr, _, err := c.doLogin(req)
	if err != nil {
		return "", err
	}

	return lr.Token, nil
}

func (c *Client) LoginWithAws(ctx context.Context, clusterName string, accessKey string, secretAccessKey string, sessionToken string, region string, profile string) (string, error) {
	var configLoadOoptions []func(*config.LoadOptions) error
	if accessKey != "" && secretAccessKey != "" {
		configLoadOoptions = append(configLoadOoptions, config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretAccessKey, sessionToken)))
		tflog.Debug(ctx, "Using credentials from provider parameters")
	}
	if region != "" {
		configLoadOoptions = append(configLoadOoptions, config.WithRegion(region))
		tflog.Debug(ctx, "Using region from provider parameters")
	}
	if profile != "" {
		configLoadOoptions = append(configLoadOoptions, config.WithSharedConfigProfile(profile))
		tflog.Debug(ctx, "Using profile from provider parameters")
	}

	cfg, err := config.LoadDefaultConfig(ctx, configLoadOoptions...)
	if err != nil {
		return "", err
	}

	tflog.Debug(ctx, fmt.Sprintf("Loaded aws config: %v", cfg))

	stsclient := sts.NewFromConfig(cfg)
	presignclient := sts.NewPresignClient(stsclient)

	stsRetriver := NewSTSTokenRetriver(presignclient)
	eksToken, err := stsRetriver.GetToken(ctx, clusterName)
	if err != nil {
		return "", err
	}

	u, err := url.JoinPath(c.address, "/auth/aws/login")
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", u, nil)
	q := req.URL.Query()
	q.Add("token", eksToken)
	req.URL.RawQuery = q.Encode()

	lr, _, err := c.doLogin(req)
	if err != nil {
		return "", err
	}

	return lr.Token, nil
}

func (c *Client) LoginWithOidc(ctx context.Context, token string, authMethod string) (string, error) {
	u, err := url.JoinPath(c.address, fmt.Sprintf("/auth/oidc/%s/login", authMethod))
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", u, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	lr, _, err := c.doLogin(req)
	if err != nil {
		return "", err
	}

	return lr.Token, nil
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
	Message   string `json:"message,omitempty"`
}
