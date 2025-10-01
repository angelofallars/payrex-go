package main

import (
	"errors"
	"fmt"

	"github.com/angelofallars/payrex-go"
)

func main() {
	payrexClient := payrex.NewClient("sk_...")

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

	paymentIntent2, err := payrexClient.PaymentIntents.Retrieve(paymentIntent.ID)
	if err != nil {
		printPayrexError(err)
		return
	}

	fmt.Printf("%+v\n", paymentIntent2)
}

// Handle errors
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
