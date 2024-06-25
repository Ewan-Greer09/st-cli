package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type HTTPClient interface {
	RegisterAgent(faction, symbol, email string) (*resty.Response, error)
}

type Client struct {
	rest *resty.Client
}

func New(baseURL string) *Client {
	resty := resty.New().SetBaseURL(baseURL)

	return &Client{
		rest: resty,
	}
}

func (c Client) RegisterAgent(faction, symbol, email string) (*resty.Response, error) {
	resp, err := c.rest.R().SetBody(struct {
		Faction string `json:"faction"`
		Symbol  string `json:"symbol"`
		Email   string `json:"email"`
	}{
		Faction: faction,
		Symbol:  symbol,
		Email:   email,
	}).Post("/register")
	if err != nil {
		return &resty.Response{}, fmt.Errorf("Register Agent: %w", err)
	}

	return resp, nil
}
