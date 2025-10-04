# PayRex-Go 🦖

[![GoDoc](https://pkg.go.dev/badge/github.com/angelofallars/payrex-go?status.svg)](https://pkg.go.dev/github.com/angelofallars/payrex-go?tab=doc)
[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/angelofallars/payrex-go/go.yml?cacheSeconds=30)](https://github.com/angelofallars/payrex-go/actions)
[![License](https://img.shields.io/github/license/angelofallars/payrex-go)](./LICENSE)
[![Stars](https://img.shields.io/github/stars/angelofallars/payrex-go)](https://github.com/angelofallars/payrex-go/stargazers)

The community-built [PayRex](https://www.payrexhq.com) Go SDK.

## Installation

Use `go get` in your Go project to install the library:

```sh
go get -u github.com/angelofallars/payrex-go
```

Then import payrex-go:

```go
import (
  "github.com/angelofallars/payrex-go"
)
```

## Getting started

For full details on the API, see the [PayRex API reference](https://docs.payrexhq.com/docs/api/core_resources).

To check out all the capabilities of this library, view the [Go package documentation](https://pkg.go.dev/github.com/angelofallars/payrex-go).

Here are a few simple examples:

### Customers

```go
payrexClient := payrex.NewClient(apiKey)
params := &payrex.CustomerCreateParams{
	Name: "Juan Dela Cruz",
	Email: "jd.cruz@gmail.com",
	Currency: payrex.CurrencyPHP,
}

customer, err := payrexClient.Customers.Create(params)
```

The API key passed in to `payrex.NewClient()` should be the Secret API Key from the PayRex dashboard which starts with `sk_`.

### PaymentIntents

```go
payrexClient := payrex.NewClient(apiKey)
params := &payrex.PaymentIntentCreateParams{
	Amount:      100_00, // represents ₱100.00
	Currency:    payrex.CurrencyPHP,
	Description: payrex.NotNil("Dino Treat"),
	PaymentMethods: payrex.Slice(
		payrex.PaymentMethodGCash,
		payrex.PaymentMethodMaya,
	),
}

paymentIntent, err := payrexClient.PaymentIntents.Create(params)
```

### Webhooks

```go
// Enable all webhooks
payrexClient := payrex.NewClient(apiKey)

webhooks, err := payrexClient.Webhooks.List(nil)
if err != nil {
	log.Fatal(err)
}

for _, webhook := range webhooks.Values {
	_, err := payrexClient.Webhooks.Enable(webhook.ID)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Webhook signing

payrex-go can verify the webhook signatures of a webhook event delivery request, and also parse the request into a `payrex.Event` value. For more info, see the [documentation for webhooks](https://docs.payrexhq.com/docs/guide/developer_handbook/webhooks).

How webhook event processing works:

```go
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
```

For more examples, see the [`payrex-go/example/`](https://github.com/angelofallars/payrex-go/tree/main/example) directory.
