package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/szoumoc/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)

	req, err := http.NewRequest("Get", endpoint, nil)
	if err != nil {
		// log.Print("ERROR")
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err == nil {
			return nil, fmt.Errorf("bad status: %s, error: %v", resp.Status, httpErr["error"])
		}
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}
	return priceResp, nil
}
