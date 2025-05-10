// Package sdk provides a client for interacting with the Bold API.
package sdk

import (
	httpClient "github.com/PChaparro/bold-co-sdk/src/internal/http"
)

// ClientConfig contains the configuration options for the BoldClient.
type ClientConfig struct {
	// ApiKey is the authentication key required to access Bold API.
	ApiKey string

	// BaseURL is the base URL for the Bold API.
	// If not provided, it defaults to "https://integrations.api.bold.co".
	BaseURL string
}

// BoldClient is a client for interacting with the Bold API.
type BoldClient struct {
	config     ClientConfig
	httpClient *httpClient.Client
}

// NewClient creates a new instance of the BoldClient.
// It requires an ApiKey and accepts an optional BaseURL.
func NewClient(config ClientConfig) *BoldClient {
	// Set default base URL if not provided
	if config.BaseURL == "" {
		config.BaseURL = "https://integrations.api.bold.co"
	}

	// Get the HTTP client instance
	client := httpClient.GetClient()

	return &BoldClient{
		config:     config,
		httpClient: client,
	}
}
