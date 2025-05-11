package definitions

// AmountType represents whether the payment amount is open or closed.
type AmountType string

const (
	AmountTypeOpen  AmountType = "OPEN"  // Payer decides the amount.
	AmountTypeClose AmountType = "CLOSE" // Merchant sets the amount.
)

// Amount represents the payment amount details
type Amount struct {
	Currency    CurrencyType `json:"currency"`        // Currency for the transaction (COP or USD, will be processed in COP).
	Taxes       []Tax        `json:"taxes,omitempty"` // Optional: List of taxes applied to the transaction (VAT or CONSUMPTION).
	TipAmount   float64      `json:"tip_amount"`      // Optional: Tip amount included in the transaction.
	TotalAmount float64      `json:"total_amount"`    // Total transaction amount including taxes and tips.
}

// CreatePaymentLinkRequest represents a payment request to the Bold API
type CreatePaymentLinkRequest struct {
	AmountType     AmountType      `json:"amount_type"`               // Required: OPEN (payer decides amount) or CLOSE (merchant sets amount).
	Amount         *Amount         `json:"amount,omitempty"`          // Required for CLOSE amount type, defines currency, taxes, tip and total amount.
	Description    string          `json:"description,omitempty"`     // Optional: Transaction description (2-100 characters).
	ExpirationDate int64           `json:"expiration_date,omitempty"` // Optional: Expiration date in Unix nanoseconds.
	CallbackURL    string          `json:"callback_url,omitempty"`    // Optional: URL to redirect after transaction (must start with https://).
	PaymentMethods []PaymentMethod `json:"payment_methods,omitempty"` // Optional: Available payment methods, if empty all methods are shown.
	PayerEmail     string          `json:"payer_email,omitempty"`     // Optional: Email to send the payment link to.
	ImageURL       string          `json:"image_url,omitempty"`       // Optional: Product image URL (must be https:// and end with .png or .jpg).
}
