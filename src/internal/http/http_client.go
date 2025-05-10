// Package http provides a singleton HTTP client for making API requests
// to external services. It simplifies making HTTP requests by providing
// methods for the most common HTTP operations (GET, POST).
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// RequestOptions contains the options for an HTTP request.
type RequestOptions struct {
	// URL to send the request to.
	URL string

	// Headers to include in the request.
	Headers map[string]string

	// Body to send with the request (for POST).
	Body interface{}

	// QueryParams to include in the URL.
	QueryParams map[string]string

	// Timeout for this specific request.
	Timeout time.Duration
}

// HTTPResponse represents the response from an HTTP .
type HTTPResponse struct {
	// StatusCode is the HTTP status code.
	StatusCode int

	// Body is the response body.
	Body []byte

	// Headers from the response.
	Headers http.Header
}

// Client is a singleton HTTP client for making requests.
type Client struct {
	httpClient *http.Client
}

var (
	instance *Client
	once     sync.Once
)

// GetClient returns the singleton instance of the HTTP client.
func GetClient() *Client {
	once.Do(func() {
		instance = &Client{
			httpClient: &http.Client{
				Timeout: 30 * time.Second, // Default timeout
			},
		}
	})
	return instance
}

// GET performs an HTTP GET request.
func (c *Client) GET(ctx context.Context, options RequestOptions) (*HTTPResponse, error) {
	// Build the URL with query parameters
	reqURL, err := c.buildURL(options.URL, options.QueryParams)
	if err != nil {
		return nil, fmt.Errorf("error building URL: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers
	c.addHeaders(req, options.Headers)

	// Set timeout for this specific request if provided
	client := c.httpClient
	if options.Timeout > 0 {
		client = &http.Client{
			Timeout: options.Timeout,
		}
	}

	// Execute request
	return c.doRequest(client, req)
}

// POST performs an HTTP POST request.
func (c *Client) POST(ctx context.Context, options RequestOptions) (*HTTPResponse, error) {
	var reqBody io.Reader

	// Process body if provided
	if options.Body != nil {
		jsonData, err := json.Marshal(options.Body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	// Build the URL with query parameters
	reqURL, err := c.buildURL(options.URL, options.QueryParams)
	if err != nil {
		return nil, fmt.Errorf("error building URL: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers
	c.addHeaders(req, options.Headers)

	// Set content type if not explicitly provided
	if reqBody != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Set timeout for this specific request if provided
	client := c.httpClient
	if options.Timeout > 0 {
		client = &http.Client{
			Timeout: options.Timeout,
		}
	}

	// Execute request
	return c.doRequest(client, req)
}

// buildURL constructs the full URL with query parameters.
func (c *Client) buildURL(baseURL string, params map[string]string) (string, error) {
	if len(params) == 0 {
		return baseURL, nil
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	query := parsedURL.Query()
	for key, value := range params {
		query.Add(key, value)
	}

	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}

// addHeaders adds the specified headers to the request.
func (c *Client) addHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// doRequest executes the HTTP request and processes the response.
func (c *Client) doRequest(client *http.Client, req *http.Request) (*HTTPResponse, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	// Close response body after reading
	defer func() {
		_ = resp.Body.Close()
	}()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Create response object
	httpResponse := &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
	}

	return httpResponse, nil
}
