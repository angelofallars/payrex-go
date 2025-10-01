// Package payrex-go provides Go applications easy access to the PayRex API.
package payrex

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const apiBaseURL = "https://api.payrexhq.com"

// Client is the main type for interacting with the PayRex API.
type Client struct {
	apiKey     string
	httpClient http.Client

	Customers      ServiceCustomers
	PaymentIntents ServicePaymentIntents
	Payments       ServicePayments
	Payouts        ServicePayouts

	// The base URL to make requests to.
	//
	// Default URL is 'https://api.payrexhq.com'
	//
	// Only override this if you know what you are doing.
	APIBaseURL string
}

// NewClient creates a new [Client] instance.
func NewClient(apiKey string) *Client {
	c := &Client{
		apiKey:     apiKey,
		httpClient: http.Client{},
		APIBaseURL: apiBaseURL,
	}

	c.loadClient()

	return c
}

// loadClient sets the unexported client field for each service implementing serviceProvider.
func (c *Client) loadClient() {
	// I use reflection here mainly just for convenience: whenever new services
	// are added I can just add them as a field to Client and I don't have to do anything else.
	v := reflect.ValueOf(c).Elem()
	for _, fieldName := range serviceFieldNames {
		service := v.FieldByName(fieldName).Addr().Interface().(serviceProvider)
		service.setupService(c)
	}
}

var serviceFieldNames []string = []string{}

// init checks that all Service fields in the Client struct implement serviceProvider.
func init() {
	v := reflect.ValueOf(&Client{}).Elem()

	for _, field := range reflect.VisibleFields(v.Type()) {
		if !field.IsExported() || !strings.HasPrefix(field.Type.Name(), "Service") {
			continue
		}

		serviceField := v.FieldByName(field.Name).Addr().Interface()

		if _, ok := serviceField.(serviceProvider); !ok {
			panic(fmt.Sprintf(
				"expected field 'Client.%s' of type '%s' to implement 'serviceProvider'",
				field.Name, field.Type.Name(),
			))
		}

		serviceFieldNames = append(serviceFieldNames, field.Name)
	}
}
