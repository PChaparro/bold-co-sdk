# SDK de Bold Colombia (CO)

<div>
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="MIT License badge" />
  <img src="https://goreportcard.com/badge/github.com/PChaparro/bold-co-sdk.svg" alt="Go report badge" />
  <img src="https://pkg.go.dev/badge/github.com/PChaparro/bold-co-sdk.svg" alt="Go reference badge" />
</div>

Este repositorio contiene un SDK **no oficial**, escrito en Go, para interactuar con la API de la pasarela de pagos [Bold](https://bold.co/).

## Funcionalidades ✨

### API de Enlaces de Pago 💸

Genera enlaces de pago de forma programática para comercios electrónicos y aplicaciones SaaS, permitiendo procesar cobros en línea de manera segura y eficiente.

- [x] Consultar métodos de pago disponibles ✅
- [x] Crear enlace de pago ✅
- [x] Consultar enlace de pago ✅

### API de Integraciones 🔌

Permite la comunicación directa entre aplicaciones y terminales de pago (datáfonos) de Bold, automatizando el proceso de cobro sin intervención manual.

- [ ] Consultar métodos de pago disponibles ❌
- [ ] Consultar terminales de pago (datáfonos) disponibles ❌
- [ ] Crear pagos ❌

## Instalación 📦

Este SDK está disponible como un paquete de Go. Puedes instalarlo como cualquier otro paquete de Go, usando el siguiente comando:

```bash
go get github.com/PChaparro/bold-co-sdk
```

## Uso 🚀

Puedes guiarte por los tests de integración para aprender a usar el SDK. A continuación, los tests disponibles por funcionalidad:

| Funcionalidad             | Tests de integración                                                              |
| ------------------------- | --------------------------------------------------------------------------------- |
| Consultar métodos de pago | [get_payment_methods_test.go](.././../../src/sdk/get_payment_methods_test.go)     |
| Crear link de pago        | [create_payment_link_test.go](.././../../src/sdk/create_payment_link_test.go)     |
| Consultar link de pago    | [get_payment_link_data_test.go](.././../../src/sdk/get_payment_link_data_test.go) |

A modo de ejemplo, para generar un enlace de pago, puedes usar el siguiente fragmento de código:

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
	// Leer la API Key de las variables de entorno
	apiKey := os.Getenv("BOLD_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Missing BOLD_API_KEY environment variable")
		os.Exit(1)
	}

	client := sdk.NewClient(sdk.ClientConfig{
		ApiKey: apiKey,
	})

	// Definir la fecha de expiración del link de pago
	expiration := time.Now().Add(24 * time.Hour).UnixNano()

	// Crear el link de pago
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
		Description:    "Mi descripción del producto o servicio",
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

  // Mostrar el enlace de pago
	fmt.Printf("Payment Link created successfully: %+v\n", response)
}
```

Este archivo incluye sólo este ejemplo para evitar redundancias. Para más casos de uso, revisa los tests de integración.

## Ejecutar pruebas 🧪

Para ejecutar las pruebas de integración, asegúrate de configurar la variable de entorno `BOLD_API_KEY` con tu clave de Bold:

```bash
export BOLD_API_KEY="your_api_key"
```

Recomendamos usar la clave del entorno de pruebas para evitar realizar peticiones a producción y prevenir transacciones no deseadas.

Luego, ejecuta las pruebas de integración con el siguiente comando:

```bash
go test -v ./src/sdk/...
```

Si deseas ver el coverage de los tests, puedes usar el siguiente comando:

```bash
# Generar reporte de cobertura de tests
go test -coverpkg=./src/sdk/... -coverprofile=./coverage/coverage.out ./src/sdk/...
# Abrir reporte de cobertura en HTML
go tool cover -html=./coverage/coverage.out
```
