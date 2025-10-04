package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/angelofallars/payrex-go"
)

func handleWebhook(webhookSecretKey string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		event, err := payrex.ParseEvent(r, webhookSecretKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}

		switch event.ResourceType {

		case payrex.EventResourceTypeBillingStatement:
			billingStatement := event.MustBillingStatement()
			prettyPrint(billingStatement)

		case payrex.EventResourceTypeCheckoutSession:
			checkoutSession := event.MustCheckoutSession()
			prettyPrint(checkoutSession)

		case payrex.EventResourceTypePaymentIntent:
			paymentIntent := event.MustPaymentIntent()
			prettyPrint(paymentIntent)

		case payrex.EventResourceTypePayout:
			payout := event.MustPayout()
			prettyPrint(payout)

		case payrex.EventResourceTypeRefund:
			refund := event.MustRefund()
			prettyPrint(refund)

		}
	})
}

func main() {
	webhookSecretKey := os.Getenv("PAYREX_WEBHOOK_SECRET")
	if webhookSecretKey == "" {
		log.Fatal("PAYREX_WEBHOOK_SECRET not set")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleWebhook(webhookSecretKey))

	server := http.Server{
		Addr:    ":2003",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func prettyPrint(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b))
}
