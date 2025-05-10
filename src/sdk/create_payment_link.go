// Package sdk provides methods to interact with Bold's payment link API
package sdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
	httpClient "github.com/PChaparro/bold-co-sdk/src/internal/http"
)

// CreatePaymentLink sends a request to create a payment link using Bold's API
// It accepts a context and a CreatePaymentLinkRequest with the necessary parameters
// Returns the API response with the payment link details or an error
func (c *BoldClient) CreatePaymentLink(ctx context.Context, req definitions.CreatePaymentLinkRequest) (*definitions.CreatePaymentLinkResponse, error) {
	// Construct the endpoint URL
	url := fmt.Sprintf("%s/online/link/v1", c.config.BaseURL)

	// Prepare headers with authentication
	headers := map[string]string{
		"Authorization": fmt.Sprintf("x-api-key %s", c.config.ApiKey),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}

	// Make the POST request
	response, err := c.httpClient.POST(ctx, httpClient.RequestOptions{
		URL:     url,
		Headers: headers,
		Body:    req,
	})

	// Handle request errors
	if err != nil {
		return nil, fmt.Errorf("failed to create payment link: %w", err)
	}

	// Check for non-successful status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("Bold API error - status code: %d, response: %s", response.StatusCode, string(response.Body))
	}

	// Parse the response
	var linkResponse definitions.CreatePaymentLinkResponse
	if err := json.Unmarshal(response.Body, &linkResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check if the response contains validation errors
	if linkResponse.HasErrors() {
		errorMessages := linkResponse.GetErrorMessages()
		// Return both the response and an error with the validation messages
		// This allows the caller to access the full response if needed
		return &linkResponse, fmt.Errorf("validation error(s): %v", errorMessages)
	}

	return &linkResponse, nil
}
