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
		Amount:      10000,
		Currency:    payrex.CurrencyPHP,
		Description: payrex.Optional("Dino Treat"),
		PaymentMethods: []payrex.PaymentMethod{
			payrex.PaymentMethodGCash,
			payrex.PaymentMethodMaya,
		},
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
