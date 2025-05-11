package sdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
	httpClient "github.com/PChaparro/bold-co-sdk/src/internal/http"
)

// GetPaymentMethodsForIntegrationsAPI retrieves the available payment methods that can be used
// with the integrations API.
func (c *BoldClient) GetPaymentMethodsForIntegrationsAPI(ctx context.Context) (*definitions.GetPaymentMethodsForIntegrationsAPIResponse, error) {
	// Construct the endpoint URL
	url := fmt.Sprintf("%s/payments/payment-methods", c.config.BaseURL)

	// Make the GET request
	response, err := c.httpClient.GET(ctx, httpClient.RequestOptions{
		URL:     url,
		Headers: httpClient.GetDefaultHeadersForBoldAPI(c.config.ApiKey),
	})

	// Handle request errors
	if err != nil {
		return nil, fmt.Errorf("failed to get available payment methods for integrations API: %w", err)
	}

	// Check for non-successful status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("bold API error - status code: %d, response: %s", response.StatusCode, string(response.Body))
	}

	// Parse the response
	var linkResponse definitions.GetPaymentMethodsForIntegrationsAPIResponse
	if err := json.Unmarshal(response.Body, &linkResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &linkResponse, nil
}
