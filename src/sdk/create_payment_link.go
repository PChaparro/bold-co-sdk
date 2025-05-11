package sdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
	httpClient "github.com/PChaparro/bold-co-sdk/src/internal/http"
)

// CreatePaymentLink sends a request to create a payment link using Bold's API.
// It accepts a context and a CreatePaymentLinkRequest with the necessary parameters.
// Returns the API response with the payment link details or an error
func (c *BoldClient) CreatePaymentLink(ctx context.Context, req definitions.CreatePaymentLinkRequest) (*definitions.CreatePaymentLinkResponse, error) {
	// Construct the endpoint URL
	url := fmt.Sprintf("%s/online/link/v1", c.config.BaseURL)

	// Make the POST request
	response, err := c.httpClient.POST(ctx, httpClient.RequestOptions{
		URL:     url,
		Headers: httpClient.GetDefaultHeadersForBoldAPI(c.config.ApiKey),
		Body:    req,
	})

	// Handle request errors
	if err != nil {
		return nil, fmt.Errorf("failed to create payment link: %w", err)
	}

	// Check for non-successful status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("bold API error - status code: %d, response: %s", response.StatusCode, string(response.Body))
	}

	// Parse the response
	var linkResponse definitions.CreatePaymentLinkResponse
	if err := json.Unmarshal(response.Body, &linkResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &linkResponse, nil
}
