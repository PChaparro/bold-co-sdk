package http

import "fmt"

// GetDefaultHeadersForBoldAPI returns the default headers for Bold API requests.
func GetDefaultHeadersForBoldAPI(apiKey string) map[string]string {
	return map[string]string{
		"Authorization": fmt.Sprintf("x-api-key %s", apiKey),
		"Accept":        "application/json",
	}
}
