package sdk

import (
	"context"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

// CreatePaymentForIntegrationsAPI sends a request to create a payment using the integrations API.
// It accepts a context and a CreatePaymentForIntegrationsAPIRequest with the necessary parameters.
// Returns the API response with the payment details or an error.
func (client *BoldClient) CreatePaymentForIntegrationsAPI(ctx context.Context, req definitions.CreatePaymentForIntegrationsAPIRequest) (*definitions.CreatePaymentForIntegrationsAPIResponse, error) {
	return sendPOSTRequest[definitions.CreatePaymentForIntegrationsAPIResponse](
		client,
		ctx,
		RequestParams{
			Endpoint: "/payments/app-checkout",
			Action:   "create payment for integrations API",
			Body:     req,
		},
	)
}
