package sdk

import (
	"context"
	"encoding/json"
	"fmt"

	httpClient "github.com/PChaparro/bold-co-sdk/src/internal/http"
)

// RequestParams encapsulates the necessary parameters to make requests to the Bold API.
type RequestParams struct {
	Endpoint string // The endpoint path, not including the baseURL.
	Action   string // Description of the action being performed (e.g., "create payment link").
	Body     any    // The request body for POST requests (optional for GET).
}

// sendGETRequest is a generic function to send GET requests to the Bold API.
// T is the type of the expected response.
func sendGETRequest[T any](
	c *BoldClient,
	ctx context.Context,
	params RequestParams,
) (*T, error) {
	// Build the complete URL.
	url := fmt.Sprintf("%s%s", c.config.BaseURL, params.Endpoint)

	// Perform the GET request.
	response, err := c.httpClient.GET(ctx, httpClient.RequestOptions{
		URL:     url,
		Headers: httpClient.GetDefaultHeadersForBoldAPI(c.config.ApiKey),
	})

	// Handle request errors.
	if err != nil {
		return nil, fmt.Errorf("failed to %s: %w", params.Action, err)
	}

	// Verify non-successful status code.
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("bold API error - status code: %d, response: %s",
			response.StatusCode, string(response.Body))
	}

	// Parse the response.
	var result T
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// sendPOSTRequest is a generic function to send POST requests to the Bold API.
// T is the type of the expected response.
func sendPOSTRequest[T any](
	c *BoldClient,
	ctx context.Context,
	params RequestParams,
) (*T, error) {
	// Build the complete URL.
	url := fmt.Sprintf("%s%s", c.config.BaseURL, params.Endpoint)

	// Perform the POST request.
	response, err := c.httpClient.POST(ctx, httpClient.RequestOptions{
		URL:     url,
		Headers: httpClient.GetDefaultHeadersForBoldAPI(c.config.ApiKey),
		Body:    params.Body,
	})

	// Handle request errors.
	if err != nil {
		return nil, fmt.Errorf("failed to %s: %w", params.Action, err)
	}

	// Verify non-successful status code.
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("bold API error - status code: %d, response: %s",
			response.StatusCode, string(response.Body))
	}

	// Parse the response.
	var result T
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}
