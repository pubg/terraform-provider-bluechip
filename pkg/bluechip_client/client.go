package bluechip_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

type Client struct {
	client  *http.Client
	token   string
	version string
	address string

	// 각 retry 사이의 대기 시간
	// 3초 권장
	retryWait time.Duration
	// Do 함수의 timeout, retry count 대신 제한된 timeout 동안 무제한 retry한다
	// 30초 권장
	timeout time.Duration
}

func NewClient(client *http.Client, token string, version string, address string, retryWait time.Duration, timeout time.Duration) *Client {
	return &Client{client: client, token: token, version: version, address: address, retryWait: retryWait, timeout: timeout}
}

func (c *Client) JoinUrl(urlPath ...string) string {
	path, err := url.JoinPath(c.address, urlPath...)
	if err != nil {
		panic(err)
	}
	return path
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("User-Agent", "bluechip-go-http/"+c.version)

	var resp *http.Response
	var err error
	err = retry.RetryContext(req.Context(), c.timeout, func() *retry.RetryError {
		resp, err = c.client.Do(req)
		if err != nil {
			time.Sleep(c.retryWait)
			return retry.RetryableError(err)
		}
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode/100 == 5 {
			time.Sleep(c.retryWait)
			return retry.NonRetryableError(fmt.Errorf("unexpected status code: %d, text: %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("response is nil")
	}

	if resp.StatusCode/100 != 2 {
		return resp, fmt.Errorf("unexpected status code: %d, text: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return resp, nil
}

func (c *Client) DoWithType(req *http.Request, t any) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return resp, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

func ReadBodyForError(resp *http.Response) []byte {
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()
	buf, _ := io.ReadAll(resp.Body)
	return buf
}
