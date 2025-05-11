package definitions

// PaymentMethodLimits represents the minimum and maximum amount limits for a payment method.
type PaymentMethodLimits struct {
	// Min is the minimum amount that can be processed with this payment method.
	Min int64 `json:"min"`

	// Max is the maximum amount that can be processed with this payment method.
	Max int64 `json:"max"`
}

// PaymentMethodsMap represents a map of available payment methods and their limits.
type PaymentMethodsMap map[PaymentMethod]PaymentMethodLimits

// PaymentMethodsData represents the available payment methods with their respective limits.
type PaymentMethodsData struct {
	// PaymentMethods contains a map of payment method types to their limits.
	PaymentMethods PaymentMethodsMap `json:"payment_methods"`
}

// GetPaymentMethodsForPaymentLinkResponse represents the response from retrieving available payment methods.
type GetPaymentMethodsForPaymentLinkResponse struct {
	// Payload contains the payment methods details.
	Payload PaymentMethodsData `json:"payload"`

	// Errors contains any errors that occurred during the request.
	Errors []ErrorField `json:"errors"`
}
