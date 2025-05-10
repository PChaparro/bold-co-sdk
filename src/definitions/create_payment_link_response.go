// Package definitions provides response structures for Bold API endpoints
package definitions

import "maps"

// PaymentLinkData represents the actual payment link information
type PaymentLinkData struct {
	// PaymentLink is the identifier of the created payment link
	PaymentLink string `json:"payment_link"`

	// URL is the payment link that can be shared with customers
	URL string `json:"url"`
}

// ErrorField represents a single error field in the API response
type ErrorField map[string]string

// CreatePaymentLinkResponse represents the response from creating a payment link
type CreatePaymentLinkResponse struct {
	// Payload contains the payment link details
	Payload struct {
		// PaymentLink is the identifier of the created payment link
		PaymentLink string `json:"payment_link"`

		// URL is the payment link that can be shared with customers
		URL string `json:"url"`
	} `json:"payload"`

	// Errors contains any errors that occurred during the request
	// Each error is a map of field name to error message
	// Example: { "amount_type": "value is not a valid enumeration member; permitted: 'CLOSE', 'OPEN'" }
	Errors []ErrorField `json:"errors"`
}

// HasErrors returns true if the response contains any errors
func (r *CreatePaymentLinkResponse) HasErrors() bool {
	return len(r.Errors) > 0
}

// GetErrorMessages returns a map of field names to error messages
func (r *CreatePaymentLinkResponse) GetErrorMessages() map[string]string {
	result := make(map[string]string)

	for _, errorField := range r.Errors {
		maps.Copy(result, errorField)
	}

	return result
}
