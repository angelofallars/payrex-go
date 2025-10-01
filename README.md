# payrex-go

The payrex-go library provides Go applications easy access to the [PayRex](https://www.payrexhq.com/) API.

## Installation

Use `go get`..

```sh
go get github.com/angelofallars/payrex-go
```

Then import payrex-go:

```go
import "github.com/angelofallars/payrex-go"
```

## Getting started

Simple usage looks like:

```go
package main

import (
	"errors"
	"fmt"

	"github.com/angelofallars/payrex-go"
)

func main() {
	payrexClient := payrex.NewClient("sk_test_...")

	paymentIntent, err := payrexClient.PaymentIntents.Retrieve("pi_...")
	if err != nil {
		printPayrexError(err)
		return
	}

	fmt.Println(paymentIntent)

	paymentIntent, err = payrexClient.PaymentIntents.Create(&payrex.CreatePaymentIntentOptions{
		Amount:         10000,
		Description:    payrex.Optional("Dino Treat"),
		PaymentMethods: []payrex.PaymentMethod{payrex.PaymentMethodGCash},
	})
	if err != nil {
		printPayrexError(err)
		return
	}

	fmt.Println(paymentIntent)
}

// Handle errors
func printPayrexError(err error) {
	var payrexError payrex.Error
	if !errors.As(err, &payrexError) {
		return
	}

	for _, errMsg := range payrexError.Errors {
		fmt.Printf("code: %v", errMsg.Code)
		fmt.Printf("detail: %v", errMsg.Detail)
		fmt.Printf("parameters: %v", errMsg.Parameter)
	}
}
```

## Additional resources

- [PayRex API Reference](https://docs.payrexhq.com/docs/api/core_resources)

## Contributing

Pull requests are always welcome!

## License

[MIT](./LICENSE)
