package client

import (
	"context"
	"encoding/json"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/vaibhavgupta04/marketdatafeeder/config"
)

type FinnhubConnector struct {
	Client *finnhub.DefaultApiService
	Ctx    context.Context
}

func NewFinnhubConnector() *FinnhubConnector {
	cfg := config.LoadFinnhubConfig()
	conf := finnhub.NewConfiguration()
	conf.AddDefaultHeader("X-Finnhub-Token", cfg.FinnhubAPI)
	client := finnhub.NewAPIClient(conf).DefaultApi
	return &FinnhubConnector{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (f *FinnhubConnector) Connect() error {
	_, _, err := f.Client.CompanyNews(f.Ctx).Symbol("AAPL").From("2020-05-01").To("2020-05-01").Execute()
	return err
}

func (f *FinnhubConnector) ReadMessage() ([]byte, error) {
	// Example: Get bid/ask for AAPL
	data, _, err := f.Client.StockBidask(f.Ctx).Symbol("AAPL").Execute()
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}

func (f *FinnhubConnector) Close() error {
	// No persistent connection to close
	return nil
}
