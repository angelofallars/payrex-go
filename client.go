// Package payrex-go provides Go applications easy access to the PayRex API.
package payrex

import (
	"net/http"
)

// Client is the main type for interacting with the PayRex API.
type Client struct {
	Customers      ServiceCustomers
	PaymentIntents ServicePaymentIntents
	Payments       ServicePayments
	Payouts        ServicePayouts
	Webhooks       ServiceWebhooks

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
