package sdk

import (
	"context"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

// GetPaymentMethodsForIntegrationsAPI retrieves the available payment methods that can be used
// with the integrations API.
func (client *BoldClient) GetPaymentMethodsForIntegrationsAPI(ctx context.Context) (*definitions.GetPaymentMethodsForIntegrationsAPIResponse, error) {
	return sendGETRequest[definitions.GetPaymentMethodsForIntegrationsAPIResponse](
		client,
		ctx,
		RequestParams{
			Endpoint: "/payments/payment-methods",
			Action:   "get available payment methods for integrations API",
		},
	)
}
