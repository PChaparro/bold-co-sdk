package sdk

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPaymentLinkData(t *testing.T) {
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

	t.Run("successful payment link data retrieval", func(t *testing.T) {
		// Create a payment link first to get a valid ID
		expirationDate := time.Now().Add(24 * time.Hour)

		req := definitions.CreatePaymentLinkRequest{
			AmountType: definitions.AmountTypeClose,
			Amount: &definitions.Amount{
				Currency: definitions.CurrencyTypeCOP,
				Taxes: []definitions.Tax{
					{
						Type:  definitions.TaxTypeIVA,
						Base:  8403,
						Value: 1597,
					},
				},
				TipAmount:   0,
				TotalAmount: 10000,
			},
			PaymentMethods: []definitions.PaymentMethod{
				definitions.PaymentMethodPse,
			},
			Description:    "Test payment link for GetPaymentLinkData",
			PayerEmail:     "test@example.com",
			ExpirationDate: expirationDate.UnixNano(),
			CallbackURL:    "https://example.com/callback",
		}

		// Create the payment link
		createResp, err := client.CreatePaymentLink(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, createResp)

		// Extract the payment link ID from the URL
		paymentLinkID := createResp.Payload.PaymentLink

		// Now get the data for this payment link
		response, err := client.GetPaymentLinkData(ctx, paymentLinkID)

		// Assert response
		require.NoError(t, err)
		require.NotNil(t, response)

		// Check fields
		assert.Equal(t, paymentLinkID, response.ID)
		assert.Equal(t, definitions.PaymentLinkStatusActive, response.Status)
		assert.Equal(t, definitions.AmountTypeClose, response.AmountType)
		assert.Equal(t, req.Description, *response.Description)
		assert.Equal(t, req.ExpirationDate, *response.ExpirationDate)
		assert.NotEmpty(t, response.CreationDate)
		assert.True(t, response.IsSandbox)

		assert.Equal(t, float64(10000), response.Total)
		assert.Equal(t, float64(8403), response.Subtotal)
		assert.Equal(t, float64(0), response.TipAmount)

		assert.Len(t, response.Taxes, 1)
		assert.Equal(t, definitions.TaxTypeIVA, response.Taxes[0].Type)
		assert.Equal(t, float64(8403), response.Taxes[0].Base)
		assert.Equal(t, float64(1597), response.Taxes[0].Value)
	})

	t.Run("non-existent payment link id", func(t *testing.T) {
		// Try to get data for a non-existent payment link ID
		nonExistentID := "LNK_NONEXISTENT"

		// Make the request
		response, err := client.GetPaymentLinkData(ctx, nonExistentID)

		// Assert error response
		require.Error(t, err)
		assert.Nil(t, response)
	})
}
