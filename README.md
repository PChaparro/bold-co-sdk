<!-- filepath: /home/pacq/Documents/Github/Personal/bold-co-sdk/README.md -->

# Bold Colombia (CO) SDK

<div>
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="MIT License badge" />
  <img src="https://goreportcard.com/badge/github.com/PChaparro/bold-co-sdk.svg" alt="Go report badge" />
  <img src="https://pkg.go.dev/badge/github.com/PChaparro/bold-co-sdk.svg" alt="Go reference badge" />
</div>

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

- [ ] Retrieve available payment methods ❌
- [ ] Retrieve available payment terminals (POS devices) ❌
- [ ] Create ❌

## Installation 📦

Install the SDK like any other Go package:

```bash
go get github.com/PChaparro/bold-co-sdk
```

## Usage 🚀

Refer to the integration tests to learn how to use the SDK. The available tests by functionality are:

| Feature             | Integration Test File                                                  |
| ------------------- | ---------------------------------------------------------------------- |
| Get payment methods | [get_payment_methods_test.go](src/sdk/get_payment_methods_test.go)     |
| Create payment link | [create_payment_link_test.go](src/sdk/create_payment_link_test.go)     |
| Get payment link    | [get_payment_link_data_test.go](src/sdk/get_payment_link_data_test.go) |

Below is an example of generating a payment link:

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
		Description:    "My product or service description",
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
