package sdk

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPaymentMethodsForIntegrationsAPI(t *testing.T) {
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

	t.Run("successful payment methods retrieval", func(t *testing.T) {
		// Make the request
		response, err := client.GetPaymentMethodsForIntegrationsAPI(ctx)

		// Assert response
		require.NoError(t, err)
		require.NotNil(t, response)

		// Check the payload
		assert.NotNil(t, response.Payload)

		// If the payment methods array is not nil, validate its structure
		if response.Payload.PaymentMethods != nil {
			assert.NotEmpty(t, *response.Payload.PaymentMethods, "Payment methods should not be empty")

			// If there are any payment methods returned, validate their structure
			for _, method := range *response.Payload.PaymentMethods {
				assert.NotEmpty(t, string(method.Name), "Payment method name should not be empty")
				assert.IsType(t, true, method.Enabled, "Enabled should be a boolean")
			}
		}

		// Check for empty errors list
		assert.Empty(t, response.Errors)
	})
}
