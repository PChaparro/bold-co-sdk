package definitions

// IntegrationPaymentMethod represents a payment method available for integrations with its enabled status.
type IntegrationPaymentMethod struct {
	// Name is the identifier of the payment method.
	Name PaymentMethod `json:"name"`

	// Enabled indicates whether the payment method is currently enabled.
	Enabled bool `json:"enabled"`
}

// IntegrationPaymentMethodsData represents the collection of available payment methods for integrations.
type IntegrationPaymentMethodsData struct {
	// PaymentMethods contains a list of available payment methods with their enabled status.
	// It can be null in the JSON response.
	PaymentMethods *[]IntegrationPaymentMethod `json:"payment_methods,omitempty"`
}

// GetPaymentMethodsForIntegrationsAPIResponse represents the response from retrieving available
// payment methods for integrations API.
type GetPaymentMethodsForIntegrationsAPIResponse struct {
	// Payload contains the payment methods details.
	Payload IntegrationPaymentMethodsData `json:"payload"`

	// Errors contains any errors that occurred during the request.
	Errors []ErrorField `json:"errors,omitempty"`
}
