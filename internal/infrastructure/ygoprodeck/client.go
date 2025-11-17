package ygoprodeck

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type Client struct {
	cfg Config
	httpClient *http.Client
	logger *zap.Logger
}

func New(ctx context.Context, cfg Config, httpClient *http.Client, logger *zap.Logger) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: httpClient,
		logger:     logger,
	}
}

func (c *Client) GetCardsRawData(ctx context.Context) ([]byte, error) {
	cardInfoEndpointURL := fmt.Sprintf("%s/cardinfo.php", c.cfg.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cardInfoEndpointURL, nil)
	if err != nil {
		// TODO: Wrap error and improve error handling
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		// TODO: Wrap error and improve error handling
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("unexpected error occured %s with status code %d", errBody, res.StatusCode)
	}
	
	b, err := io.ReadAll(res.Body)
	if err != nil {
		// TODO: Wrap error and improve error handling
		return nil, err
	}

	return b, nil
}