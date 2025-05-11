package sdk

import (
	"context"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

// GetBindedTerminalsForIntegrationsAPI retrieves the binded terminals
// that can be used with the integrations API.
func (client *BoldClient) GetBindedTerminalsForIntegrationsAPI(ctx context.Context) (*definitions.GetBindedTerminalsForIntegrationsAPIResponse, error) {
	return sendGETRequest[definitions.GetBindedTerminalsForIntegrationsAPIResponse](
		client,
		ctx,
		RequestParams{
			Endpoint: "/payments/binded-terminals",
			Action:   "get binded terminals for integrations API",
		},
	)
}
