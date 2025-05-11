package sdk

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPaymentMethodsForPaymentLink(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("successful payment methods retrieval", func(t *testing.T) {
		// Make the request
		response, err := client.GetPaymentMethodsForPaymentLink(ctx)

		// Assert response
		require.NoError(t, err)
		require.NotNil(t, response)

		// Check the payload
		assert.NotNil(t, response.Payload)
		assert.NotNil(t, response.Payload.PaymentMethods)

		// Verify that we have at least one payment method
		assert.Greater(t, len(response.Payload.PaymentMethods), 0, "Should have at least one payment method")

		// Verify that if methods are returned, they have min and max values
		methods := response.Payload.PaymentMethods

		for _, limits := range methods {
			assert.NotZero(t, limits.Min, "Min limit should be set")
			assert.NotZero(t, limits.Max, "Max limit should be set")
		}

		// Check for empty errors list
		assert.Empty(t, response.Errors)
	})
}
