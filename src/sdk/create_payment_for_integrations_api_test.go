package sdk

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/PChaparro/bold-co-sdk/src/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentForIntegrationsAPI(t *testing.T) {
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

	t.Run("expecting terminal not available error", func(t *testing.T) {
		// Create a payment request using our helper
		req := tests.GetPayloadToCreateValidPaymentForIntegrationsAPI()

		// Make the request
		response, err := client.CreatePaymentForIntegrationsAPI(ctx, *req)

		// We expect an error about terminal not available
		require.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "Terminal no disponible"),
			"Expected error to contain 'Terminal no disponible', got: %v", err)
		assert.Nil(t, response)
	})
}
