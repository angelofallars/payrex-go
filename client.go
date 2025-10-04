// Package payrex-go provides Go applications easy access to the PayRex API.
package payrex

import (
	"net/http"
)

// Client is the PayRex client. It contains all the services
// for interacting with the PayRex API.
type Client struct {
	// BillingStatementLineItems is the service for invoking /billing_statement_line_items APIs.
	BillingStatementLineItems ServiceBillingStatementLineItems
	// BillingStatements is the service for invoking /billing_statements APIs.
	BillingStatements ServiceBillingStatements
	// CheckoutSessions is the service for invoking /checkout_sessions APIs.
	CheckoutSessions ServiceCheckoutSessions
	// CustomerSessions is the service for invoking /customer_sessions APIs.
	CustomerSessions ServiceCustomerSessions
	// Customers is the service for invoking /customers APIs.
	Customers ServiceCustomers
	// PaymentIntents is the service for invoking /payment_intents APIs.
	PaymentIntents ServicePaymentIntents
	// Payments is the service for invoking /payments APIs.
	Payments ServicePayments
	// Payouts is the service for invoking /payouts APIs.
	Payouts ServicePayouts
	// Refunds is the service for invoking /refunds APIs.
	Refunds ServiceRefunds
	// Webhooks is the service for invoking /webhooks APIs.
	Webhooks ServiceWebhooks

	apiBaseURL string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new [Client] instance.
func NewClient(apiKey string) *Client {
	const apiBaseURL = "https://api.payrexhq.com"

	c := &Client{
		apiBaseURL: apiBaseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}

	c.setupServices()

	return c
}

// WithHTTPClient replaces the default HTTP client used for making requests.
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient
	return c
}
