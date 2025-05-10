package sdk

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
	"github.com/PChaparro/bold-co-sdk/src/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentLink(t *testing.T) {
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

	t.Run("successful payment link creation", func(t *testing.T) {
		// Create a payment link request using our helper
		req := tests.GetPayloadToCreateValidPaymentLink()

		// Make the request
		response, err := client.CreatePaymentLink(ctx, *req)

		// Assert response
		require.NoError(t, err)
		require.NotNil(t, response)
		assert.NotEmpty(t, response.Payload.PaymentLink)
		assert.NotEmpty(t, response.Payload.URL)
		assert.Empty(t, response.Errors)
	})

	t.Run("with missing required fields", func(t *testing.T) {
		// Prepare request without required fields
		req := definitions.CreatePaymentLinkRequest{
			Description: "Test payment link",
		}

		// Make the request
		response, err := client.CreatePaymentLink(ctx, req)

		// Assert error response
		require.Error(t, err)
		require.Nil(t, response)
	})
}
