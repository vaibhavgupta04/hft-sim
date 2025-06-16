package client

import (
	"fmt"
	"io"
	"net/http"
	"github.com/vaibhavgupta04/marketdatafeeder/config"
)

type MarketStackConnector struct {
	APIKey string
	URL    string
	Client *http.Client
}

func NewMarketStackConnector() *MarketStackConnector {
	cfg := config.LoadMarketstackConfig()
	return &MarketStackConnector{
		APIKey: cfg.MarketstackAPI,
		URL:    cfg.MarketstackURL,
		Client: &http.Client{},
	}
}

func (m *MarketStackConnector) Connect() error {
	// REST-based, so nothing to connect. Maybe test endpoint.
	return nil
}

func (m *MarketStackConnector) ReadMessage() ([]byte, error) {
	url := fmt.Sprintf("%s/tickers?access_key=%s", m.URL, m.APIKey)
	resp, err := m.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (m *MarketStackConnector) Close() error {
	return nil
}

type MarketstackClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// Example method to fetch data (you can expand as needed)
func (c *MarketstackClient) Get(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("access_key", c.APIKey)
	req.URL.RawQuery = q.Encode()
	return c.HTTPClient.Do(req)
}
