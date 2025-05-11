package sdk

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBindedTerminalsForIntegrationsAPI(t *testing.T) {
	// Get API key from environment variable
	apiKey := os.Getenv("BOLD_API_KEY")
	if apiKey == "" {
		t.Skip("BOLD_API_KEY environment variable not set")
	}

	// Initialize the client
	client := NewClient(ClientConfig{
		ApiKey: apiKey,
	})

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Due to the absence of a payment terminal, the API does not return
	// an empty array but instead returns a 404 error, indicating that no
	// payment terminals were found. If a payment terminal were available,
	// the following code should be used to validate the response.
	/* t.Run("successful binded terminals retrieval", func(t *testing.T) {
		// Make the request
		response, err := client.GetBindedTerminalsForIntegrationsAPI(ctx)

		// Assert response
		require.NoError(t, err)
		require.NotNil(t, response)

		// Check the payload
		assert.NotNil(t, response.Payload)

		// If the available terminals array is not nil, validate its structure
		if response.Payload.AvailableTerminals != nil {
			assert.NotEmpty(t, *response.Payload.AvailableTerminals, "Available terminals should not be empty")

			// If there are any terminals returned, validate their structure
			for _, terminal := range *response.Payload.AvailableTerminals {
				assert.NotEmpty(t, terminal.TerminalModel, "Terminal model should not be empty")
				assert.NotEmpty(t, terminal.TerminalSerial, "Terminal serial should not be empty")
				assert.NotEmpty(t, terminal.Name, "Terminal name should not be empty")
				assert.NotEmpty(t, string(terminal.Status), "Terminal status should not be empty")
			}
		}

		// Check for empty errors list
		assert.Empty(t, response.Errors)
	}) */

	// Workaround for the absence of a payment .
	// The API does not return an empty array but instead returns a 404 error,
	t.Run("available terminals not found", func(t *testing.T) {
		// Make the request
		response, err := client.GetBindedTerminalsForIntegrationsAPI(ctx)

		// Assert response
		require.Error(t, err)
		require.Nil(t, response)
		assert.Contains(t, err.Error(), "Available terminals not found", "Error message should contain 'Available terminals not found'")
	})
}
