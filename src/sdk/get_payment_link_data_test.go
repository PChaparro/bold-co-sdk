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
		// Create a payment link request using our helper
		createPaymentLinkRequest := tests.GetPayloadToCreateValidPaymentLink()

		// Create the payment link
		createPaymentLinkResponse, _ := client.CreatePaymentLink(ctx, *createPaymentLinkRequest)
		require.NotEmpty(t, createPaymentLinkResponse.Payload.PaymentLink)

		// Extract the payment link ID from the response
		paymentLinkID := createPaymentLinkResponse.Payload.PaymentLink

		// Now get the data for this payment link
		paymentLinkDataResponse, err := client.GetPaymentLinkData(ctx, paymentLinkID)

		// Assert response
		require.NoError(t, err)
		require.NotNil(t, paymentLinkDataResponse)

		// Check fields
		assert.Equal(t, paymentLinkID, paymentLinkDataResponse.ID)
		assert.Equal(t, definitions.PaymentLinkStatusActive, paymentLinkDataResponse.Status)
		assert.Equal(t, definitions.AmountTypeClose, paymentLinkDataResponse.AmountType)
		assert.Equal(t, createPaymentLinkRequest.Description, *paymentLinkDataResponse.Description)
		assert.Equal(t, createPaymentLinkRequest.ExpirationDate, *paymentLinkDataResponse.ExpirationDate)
		assert.NotEmpty(t, paymentLinkDataResponse.CreationDate)
		assert.True(t, paymentLinkDataResponse.IsSandbox)

		assert.Equal(t, createPaymentLinkRequest.Amount.TotalAmount, paymentLinkDataResponse.Total)
		assert.Equal(t, createPaymentLinkRequest.Amount.TipAmount, paymentLinkDataResponse.TipAmount)
		expectedSubtotal := createPaymentLinkRequest.Amount.TotalAmount - createPaymentLinkRequest.Amount.TipAmount - createPaymentLinkRequest.Amount.Taxes[0].Value
		assert.Equal(t, expectedSubtotal, paymentLinkDataResponse.Subtotal)

		assert.Len(t, paymentLinkDataResponse.Taxes, len(createPaymentLinkRequest.Amount.Taxes))
		for i, tax := range paymentLinkDataResponse.Taxes {
			assert.Equal(t, createPaymentLinkRequest.Amount.Taxes[i].Type, tax.Type)
			assert.Equal(t, createPaymentLinkRequest.Amount.Taxes[i].Base, tax.Base)
			assert.Equal(t, createPaymentLinkRequest.Amount.Taxes[i].Value, tax.Value)
		}
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
