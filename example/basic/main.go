package main

import (
	"fmt"
	"log"
	"os"

	"github.com/angelofallars/payrex-go"
)

func main() {
	payrexClient := payrex.NewClient(os.Getenv("PAYREX_API_KEY"))

	// Create a PaymentIntent
	paymentIntent, err := payrexClient.PaymentIntents.Create(&payrex.PaymentIntentCreateParams{
		Amount:      100_00, // represents ₱100.00
		Currency:    payrex.CurrencyPHP,
		Description: payrex.NotNil("Dino Treat"),
		PaymentMethods: payrex.Slice(
			payrex.PaymentMethodGCash,
			payrex.PaymentMethodMaya,
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve a PaymentIntent
	paymentIntent, err = payrexClient.PaymentIntents.Retrieve(paymentIntent.ID)
	if err != nil {
		log.Fatal(err)
	}

	// Cancel a PaymentIntent
	_, err = payrexClient.PaymentIntents.Cancel(paymentIntent.ID)
	if err != nil {
		log.Fatal(err)
	}

	// List customers
	customers, err := payrexClient.Customers.List(nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, customer := range customers.Values {
		fmt.Println(customer.Name)
	}
}
