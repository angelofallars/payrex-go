# payrex-go 🦖

[![GoDoc](https://pkg.go.dev/badge/github.com/angelofallars/payrex-go?status.svg)](https://pkg.go.dev/github.com/angelofallars/payrex-go?tab=doc)
[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/angelofallars/payrex-go/go.yml?cacheSeconds=30)](https://github.com/angelofallars/payrex-go/actions)
[![License](https://img.shields.io/github/license/angelofallars/payrex-go)](./LICENSE)
[![Stars](https://img.shields.io/github/stars/angelofallars/payrex-go)](https://github.com/angelofallars/payrex-go/stargazers)

Gain easy powerful access to the [PayRex](https://www.payrexhq.com/) API in Go with the payrex-go SDK.

This library is designed to have a similar interface to the [official PayRex SDKs](https://docs.payrexhq.com/docs/guide/developer_handbook/libraries_and_tools) for familiarity.

## Installation

Use `go get` in your Go project to install the library:

```sh
go get github.com/angelofallars/payrex-go
```

Then import the library:

```go
import (
  "github.com/angelofallars/payrex-go"
)
```

## Getting started

Basic usage looks like this:

```go
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/angelofallars/payrex-go"
)

func main() {
	payrexClient := payrex.NewClient(os.Getenv("PAYREX_API_KEY"))

	paymentIntent, err := payrexClient.PaymentIntents.Create(&payrex.CreatePaymentIntentOptions{
		Amount:         10000,
		Currency:       payrex.CurrencyPHP,
		Description:    payrex.Optional("Dino Treat"),
		PaymentMethods: []payrex.PaymentMethod{payrex.PaymentMethodGCash},
	})
	if err != nil {
		printPayrexError(err)
		return
	}
	fmt.Printf("%+v\n", paymentIntent)

	paymentIntent, err = payrexClient.PaymentIntents.Retrieve(paymentIntent.ID)
	if err != nil {
		printPayrexError(err)
		return
	}
	fmt.Printf("%+v\n", paymentIntent)

	paymentIntent, err = payrexClient.PaymentIntents.Cancel(paymentIntent.ID)
	if err != nil {
		printPayrexError(err)
		return
	}
	fmt.Printf("%+v\n", paymentIntent)
}

func printPayrexError(err error) {
	var payrexError payrex.Error
	if !errors.As(err, &payrexError) {
		return
	}

	for _, errMsg := range payrexError.Errors {
		fmt.Printf("code: %v\n", errMsg.Code)
		fmt.Printf("detail: %v\n", errMsg.Detail)
		fmt.Printf("parameters: %v\n", errMsg.Parameter)
	}
}
```

### WIP progress

Progress on implemented resources:
- [ ] Billing Statement Line Items
- [ ] Billing Statements
- [ ] Checkout Sessions
- [ ] Customer Sessions
- [x] Customers
- [x] Payment Intents
- [x] Payments
- [x] Payouts
- [ ] Refunds
- [ ] Webhooks

## Additional resources

- [PayRex API Reference](https://docs.payrexhq.com/docs/api/core_resources)

## Contributing

Pull requests are always welcome!

## License

[MIT](./LICENSE)
