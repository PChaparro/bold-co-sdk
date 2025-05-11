package definitions

// IntegrationPaymentData represents the information of a created integration payment.
type IntegrationPaymentData struct {
	// IntegrationID is the unique identifier of the created payment integration.
	IntegrationID string `json:"integration_id"`
}

// CreatePaymentForIntegrationsAPIResponse represents the response from creating
// a payment using the integrations API.
type CreatePaymentForIntegrationsAPIResponse struct {
	// Payload contains the payment integration details.
	Payload IntegrationPaymentData `json:"payload"`

	// Errors contains any errors that occurred during the request.
	Errors []ErrorField `json:"errors,omitempty"`
}
