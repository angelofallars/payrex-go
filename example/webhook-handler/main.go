package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/angelofallars/payrex-go"
)

var webhookSecretKey = os.Getenv("PAYREX_WEBHOOK_SECRET")

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	event, err := payrex.ParseEvent(r, webhookSecretKey)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	switch event.ResourceType {
	case payrex.EventResourceTypeBillingStatement:
		billingStatement := event.MustBillingStatement()
		fmt.Printf("%+v\n", billingStatement)

	case payrex.EventResourceTypeCheckoutSession:
		checkoutSession := event.MustCheckoutSession()
		fmt.Printf("%+v\n", checkoutSession)

	case payrex.EventResourceTypePaymentIntent:
		paymentIntent := event.MustPaymentIntent()
		fmt.Printf("%+v\n", paymentIntent)

	case payrex.EventResourceTypePayout:
		payout := event.MustPayout()
		fmt.Printf("%+v\n", payout)

	case payrex.EventResourceTypeRefund:
		refund := event.MustRefund()
		fmt.Printf("%+v\n", refund)
	}
}

func main() {
	if webhookSecretKey == "" {
		log.Fatal("PAYREX_WEBHOOK_SECRET not set")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleWebhook)

	server := http.Server{
		Addr:    ":2003",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
