package sdk

import (
	"context"
	"fmt"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

func (client *BoldClient) GetPaymentLinkData(
	ctx context.Context,
	paymentLinkId string,
) (*definitions.GetPaymentLinkDataResponse, error) {
	return sendGETRequest[definitions.GetPaymentLinkDataResponse](
		client,
		ctx,
		RequestParams{
			Endpoint: fmt.Sprintf("/online/link/v1/%s", paymentLinkId),
			Action:   "get data of payment link",
		},
	)
}
