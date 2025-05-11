package definitions

// PaymentLinkStatus represents the current status of a payment link.
type PaymentLinkStatus string

const (
	PaymentLinkStatusActive     PaymentLinkStatus = "ACTIVE"     // The payment link is active and can be used for payment.
	PaymentLinkStatusProcessing PaymentLinkStatus = "PROCESSING" // The payment is being processed.
	PaymentLinkStatusPaid       PaymentLinkStatus = "PAID"       // The payment link has been paid.
	PaymentLinkStatusRejected   PaymentLinkStatus = "REJECTED"   // The payment was rejected.
	PaymentLinkStatusExpired    PaymentLinkStatus = "EXPIRED"    // The payment link has expired.
	PaymentLinkStatusCanceled   PaymentLinkStatus = "CANCELED"   // The payment link has been canceled.
)

// PaymentLinkDetails represents the detailed information of a payment link.
type PaymentLinkDetails struct {
	// APIVersion is the version of the API that was used to create this payment link.
	APIVersion int `json:"api_version"`

	// ID is the unique identifier for the payment link.
	ID string `json:"id"`

	// Total is the total amount to be paid.
	Total float64 `json:"total"`

	// Subtotal is the amount before taxes and tips.
	Subtotal float64 `json:"subtotal"`

	// TipAmount is the optional tip amount included in the payment.
	TipAmount float64 `json:"tip_amount"`

	// Taxes is a list of taxes applied to the payment.
	Taxes []Tax `json:"taxes"`

	// Status represents the current status of the payment link.
	Status PaymentLinkStatus `json:"status"`

	// ExpirationDate is the date when the payment link will expire.
	// It can be null if the link doesn't have an expiration date.
	ExpirationDate *int64 `json:"expiration_date,omitempty"`

	// CreationDate is the timestamp when the payment link was created.
	CreationDate int64 `json:"creation_date"`

	// Description is an optional description of the payment.
	// It can be null if no description was provided.
	Description *string `json:"description,omitempty"`

	// PaymentMethod is the method used for payment.
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`

	// TransactionID is the ID of the transaction if the payment has been completed.
	// It can be null if the payment hasn't been processed yet.
	TransactionID *string `json:"transaction_id,omitempty"`

	// AmountType indicates whether the amount is open (payer decides) or
	// close (merchant sets amount).
	AmountType AmountType `json:"amount_type"`

	// IsSandbox indicates whether the payment link was created in sandbox mode.
	IsSandbox bool `json:"is_sandbox"`
}

// GetPaymentLinkDataResponse represents the response from retrieving payment link information.
// For this specific endpoint, the API returns the data directly without wrapping it in a payload.
type GetPaymentLinkDataResponse struct {
	// Embed all fields from PaymentLinkDetails directly in this struct
	PaymentLinkDetails

	// Errors contains any errors that occurred during the request.
	Errors []ErrorField `json:"errors,omitempty"`
}
