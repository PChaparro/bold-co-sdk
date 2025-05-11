package sdk

import (
	"context"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

// GetPaymentMethodsForPaymentLink retrieves the available payment methods that can be used
// for creating a payment link.
func (client *BoldClient) GetPaymentMethodsForPaymentLink(ctx context.Context) (*definitions.GetPaymentMethodsForPaymentLinkResponse, error) {
	return sendGETRequest[definitions.GetPaymentMethodsForPaymentLinkResponse](
		client,
		ctx,
		RequestParams{
			Endpoint: "/online/link/v1/payment_methods",
			Action:   "get available payment methods for payment link",
		},
	)
}
