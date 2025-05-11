# Bold Colombia (CO) SDK

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Go Report Card](https://goreportcard.com/badge/github.com/PChaparro/bold-co-sdk)](https://goreportcard.com/report/github.com/PChaparro/bold-co-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/PChaparro/bold-co-sdk.svg)](https://pkg.go.dev/github.com/PChaparro/bold-co-sdk)

This repository contains an **unofficial** Go SDK for interacting with the [Bold](https://bold.co/) payment gateway API.

## Translations 🌐

[English](README.md) | [Español](docs/i18n/es/README.md)

## Features ✨

### Payment Links API 💸

Programmatically generate payment links for e-commerce and SaaS platforms, enabling secure and efficient online transactions.

- [x] Retrieve available payment methods ✅
- [x] Create payment link ✅
- [x] Retrieve payment link ✅

### Integrations API 🔌

Facilitates direct communication between applications and Bold payment terminals (point-of-sale devices), automating the checkout process without manual intervention.

- [x] Retrieve available payment methods ✅ (Beta, **not fully tested**)
- [x] Retrieve available payment terminals (POS devices) ✅ (Beta, **not fully tested**)
- [ ] Create ❌

Please note that the **integrations API is currently in beta**, which means it may undergo changes. Likewise, the implementation in this repository for interacting with the integrations API is also in beta, as testing these endpoints requires a physical payment terminal. **The current implementation is based solely on the examples provided in the official documentation, and full testing has not been possible due to the lack of access to a physical device.**

## Installation 📦

Install the SDK like any other Go package:

```bash
go get github.com/PChaparro/bold-co-sdk
```

## Usage 🚀

Refer to the integration tests to learn how to use the SDK. The available tests by functionality are:

| Feature                               | Integration Test File                                                                                          |
| ------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Get payment methods for payment link  | [get_payment_methods_for_payment_link_test.go](src/sdk/get_payment_methods_for_payment_link_test.go)           |
| Create payment link                   | [create_payment_link_test.go](src/sdk/create_payment_link_test.go)                                             |
| Get payment link                      | [get_payment_link_data_test.go](src/sdk/get_payment_link_data_test.go)                                         |
| Get payment methods for integration   | [get_payment_methods_for_integrations_api_test.go](src/sdk/get_payment_methods_for_integrations_api_test.go)   |
| Get payment terminals for integration | [get_binded_terminals_for_integrations_api_test.go](src/sdk/get_binded_terminals_for_integrations_api_test.go) |

Below is an example of generating a payment link:

<details>
<summary>Click to expand</summary>

```go
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/PChaparro/bold-co-sdk/src/definitions"
	"github.com/PChaparro/bold-co-sdk/src/sdk"
)

func main() {
	// Load the API key from environment variables
	apiKey := os.Getenv("BOLD_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Missing BOLD_API_KEY environment variable")
		os.Exit(1)
	}

	client := sdk.NewClient(sdk.ClientConfig{
		ApiKey: apiKey,
	})

	// Set the payment link expiration date
	expiration := time.Now().Add(24 * time.Hour).UnixNano()

	// Create the payment link request
	paymentLinkRequest := definitions.CreatePaymentLinkRequest{
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
		PayerEmail:     "johndoe@example.com",
		ImageURL:       "https://robohash.org/sad.png",
		ExpirationDate: expiration,
		CallbackURL:    "https://example.com/callback",
	}

	ctx := context.Background()
	response, err := client.CreatePaymentLink(ctx, paymentLinkRequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating payment link: %v\n", err)
		os.Exit(1)
	}

	// Print the payment link
	fmt.Printf("Payment link created successfully: %+v\n", response)
}
```

</details>

## Running Tests 🧪

Ensure the `BOLD_API_KEY` environment variable is set with your Bold API key:

```bash
export BOLD_API_KEY="your_api_key"
```

We recommend using your sandbox API key to avoid executing real transactions in production.

Run the integration tests:

```bash
go test -v ./src/sdk/...
```

To generate a test coverage report:

```bash
# Generate coverage report
go test -coverpkg=./src/sdk/... -coverprofile=./coverage/coverage.out ./src/sdk/...
# Open coverage report in HTML
go tool cover -html=./coverage/coverage.out
```
