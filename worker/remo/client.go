package remo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Client is Nature Remo API client
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// NewClient returns new client with access token
func NewClient(token string) *Client {
	return &Client{"https://api.nature.global/1", http.DefaultClient, token}
}

// GetDevices return remo devices with newest sensor values
func (c *Client) GetDevices(ctx context.Context) ([]*Device, error) {
	var devices []*Device

	url := c.baseURL + "/devices"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, c.error(resp.StatusCode, resp.Body)
		// return nil, fmt.Errorf("request failed: %s", resp.StatusCode) // TODO: add body
	}

	if err := json.NewDecoder(resp.Body).Decode(&devices); err != nil {
		return nil, fmt.Errorf("invalid response body: %s", err)
	}
	return devices, nil
}

func (c *Client) error(statusCode int, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil || len(buf) == 0 {
		return fmt.Errorf("request failed with status code %d", statusCode)
	}
	return fmt.Errorf("StatusCode: %d, Error: %s", statusCode, string(buf))
}
