package definitions

// IntegrationAmount represents the payment amount details for the integration API
type IntegrationAmount struct {
	Currency    CurrencyType `json:"currency"`        // Currency for the transaction (e.g., COP)
	Taxes       []Tax        `json:"taxes,omitempty"` // List of taxes applied to the transaction
	TipAmount   float64      `json:"tip_amount"`      // Optional: Tip amount included in the transaction
	TotalAmount float64      `json:"total_amount"`    // Total transaction amount including taxes and tips
}

// PayerDocument represents the identification document of the payer for integration API
type IntegrationPayerDocument struct {
	DocumentType   DocumentType `json:"document_type"`   // Type of ID document (e.g., CEDULA, NIT, etc.)
	DocumentNumber string       `json:"document_number"` // ID document number (4-15 characters)
}

// IntegrationPayer represents the person making the payment in integration API
type IntegrationPayer struct {
	Email       string                    `json:"email,omitempty"`        // Payer's email address
	PhoneNumber string                    `json:"phone_number,omitempty"` // Payer's phone number
	Document    *IntegrationPayerDocument `json:"document,omitempty"`     // Payer's identification document
}

// CreatePaymentForIntegrationsAPIRequest represents a payment request to the Bold Integration API
type CreatePaymentForIntegrationsAPIRequest struct {
	Amount         IntegrationAmount `json:"amount"`                // Required: Contains detailed information about the amounts to be processed
	UserEmail      string            `json:"user_email"`            // Required: Email of the person making the sale
	PaymentMethod  PaymentMethod     `json:"payment_method"`        // Required: Payment method (POS, NEQUI, DAVIPLATA, PAY_BY_LINK, or empty string)
	TerminalModel  string            `json:"terminal_model"`        // Required: Model of the terminal to process the payment
	TerminalSerial string            `json:"terminal_serial"`       // Required: Serial of the terminal to process the payment
	Reference      string            `json:"reference"`             // Required: Reference to identify the payment in the Webhook response
	Description    string            `json:"description,omitempty"` // Optional: Brief description of the transaction
	Payer          *IntegrationPayer `json:"payer,omitempty"`       // Optional: Object specifying the payer's details if needed
}
