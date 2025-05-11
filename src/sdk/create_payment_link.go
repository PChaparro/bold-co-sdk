package sdk

import (
	"context"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

// CreatePaymentLink sends a request to create a payment link using Bold's API.
// It accepts a context and a CreatePaymentLinkRequest with the necessary parameters.
// Returns the API response with the payment link details or an error
func (client *BoldClient) CreatePaymentLink(ctx context.Context, req definitions.CreatePaymentLinkRequest) (*definitions.CreatePaymentLinkResponse, error) {
	return sendPOSTRequest[definitions.CreatePaymentLinkResponse](
		client,
		ctx,
		RequestParams{
			Endpoint: "/online/link/v1",
			Action:   "create payment link",
			Body:     req,
		},
	)
}
