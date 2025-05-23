package definitions

// PaymentLinkData represents the actual payment link information.
type PaymentLinkData struct {
	// PaymentLink is the identifier of the created payment link.
	PaymentLink string `json:"payment_link"`

	// URL is the payment link that can be shared with customers.
	URL string `json:"url"`
}

// CreatePaymentLinkResponse represents the response from creating a payment link.
type CreatePaymentLinkResponse struct {
	// Payload contains the payment link details.
	Payload PaymentLinkData `json:"payload"`

	// Errors contains any errors that occurred during the request.
	Errors []ErrorField `json:"errors"`
}
