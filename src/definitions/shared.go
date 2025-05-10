// Package definitions contains all the shared structures and data types
// used to build requests to the different Bold API endpoints.
// These types are common and can be reused across various operations with the API.
package definitions

// DocumentType represents the supported document types for identification
type DocumentType string

const (
	DocumentTypeCedula               DocumentType = "CEDULA"
	DocumentTypeNit                  DocumentType = "NIT"
	DocumentTypeCedulaExtranjeria    DocumentType = "CEDULA_EXTRANJERIA"
	DocumentTypePep                  DocumentType = "PEP"
	DocumentTypePasaporte            DocumentType = "PASAPORTE"
	DocumentTypeNuip                 DocumentType = "NUIP"
	DocumentTypeRegistroCivil        DocumentType = "REGISTRO_CIVIL"
	DocumentTypeDocumentoExtranjeria DocumentType = "DOCUMENTO_EXTRANJERIA"
	DocumentTypeTarjetaIdentidad     DocumentType = "TARJETA_IDENTIDAD"
	DocumentTypePpt                  DocumentType = "PPT"
)

// CurrencyType represents the supported currency types
type CurrencyType string

const (
	CurrencyTypeCOP CurrencyType = "COP" // Colombian Peso
)

// PaymentMethod represents the supported payment methods
type PaymentMethod string

const (
	PaymentMethodCreditCard       PaymentMethod = "CREDIT_CARD"
	PaymentMethodPse              PaymentMethod = "PSE"
	PaymentMethodBotonBancolombia PaymentMethod = "BOTON_BANCOLOMBIA"
	PaymentMethodNequi            PaymentMethod = "NEQUI"
)

// PayerDocument represents the identification document of the payer
type PayerDocument struct {
	DocumentType   DocumentType `json:"document_type"`
	DocumentNumber string       `json:"document_number"`
}

// Payer represents the person making the payment
type Payer struct {
	Email       string        `json:"email,omitempty"`
	PhoneNumber string        `json:"phone_number,omitempty"`
	Document    PayerDocument `json:"document"`
}

// TaxType represents the type of tax applied to the payment
type TaxType string

const (
	TaxTypeIVA TaxType = "VAT" // Value Added Tax
)

// Tax represents a tax applied to the payment
type Tax struct {
	Type  TaxType `json:"type"` // Currently only "VAT" is supported
	Base  float64 `json:"base"`
	Value float64 `json:"value"`
}
