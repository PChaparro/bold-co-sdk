package tests

import (
	"time"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
)

// GetPayloadToCreateValidPaymentLink is a helper function that returns a valid
// payload for creating a payment link. It is only intended for testing purposes
// to reduce code duplication in test cases.
func GetPayloadToCreateValidPaymentLink() *definitions.CreatePaymentLinkRequest {
	// Create a payment link request
	expirationDate := time.Now().Add(1 * time.Minute)

	req := &definitions.CreatePaymentLinkRequest{
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
		Description:    "Description of product or service",
		ExpirationDate: expirationDate.UnixNano(),
		CallbackURL:    "https://example.com/callback",
		ImageURL:       "https://robohash.org/sad.png",
		PayerEmail:     "johndoe@example.com",
	}

	return req
}
