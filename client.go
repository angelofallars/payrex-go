// Package payrex-go provides Go applications easy access to the PayRex API.
package payrex

import (
	"net/http"
)

// Client manages interactions with the PayRex API.
type Client struct {
	// Services used for interacting with the different parts of the PayRex API.
	CheckoutSessions ServiceCheckoutSessions
	Customers        ServiceCustomers
	PaymentIntents   ServicePaymentIntents
	Payments         ServicePayments
	Payouts          ServicePayouts
	Refunds          ServiceRefunds
	Webhooks         ServiceWebhooks

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
