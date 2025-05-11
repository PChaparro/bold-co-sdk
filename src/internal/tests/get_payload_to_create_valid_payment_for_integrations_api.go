package tests

import (
	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

// GetPayloadToCreateValidPaymentForIntegrationsAPI returns a valid request to create a payment
// using the Bold integrations API for testing purposes.
func GetPayloadToCreateValidPaymentForIntegrationsAPI() *definitions.CreatePaymentForIntegrationsAPIRequest {
	return &definitions.CreatePaymentForIntegrationsAPIRequest{
		Amount: definitions.IntegrationAmount{
			Currency: definitions.CurrencyTypeCOP,
			Taxes: []definitions.Tax{
				{
					Type:  definitions.TaxTypeIVA,
					Base:  10000,
					Value: 1000,
				},
			},
			TipAmount:   0,
			TotalAmount: 1230000,
		},
		UserEmail:      "seller@merchant.com",
		PaymentMethod:  definitions.PaymentMethodPos,
		TerminalModel:  "N86",
		TerminalSerial: "N860W000000",
		Reference:      "d9b10690-981d-494d-bcb0-66a1dacab51d",
		Description:    "Test Purchase",
		Payer: &definitions.IntegrationPayer{
			Email:       "customer@example.com",
			PhoneNumber: "3100000000",
			Document: &definitions.IntegrationPayerDocument{
				DocumentType:   definitions.DocumentTypeCedula,
				DocumentNumber: "1010140000",
			},
		},
	}
}
