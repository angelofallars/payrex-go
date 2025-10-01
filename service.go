package payrex

import (
	"fmt"
	"reflect"
	"strings"
)

// service is the base type that all Service types embed.
type service struct {
	client *Client
}

func (s *service) setupService(client *Client) {
	s.client = client
}

// serviceProvider is the interface that all Service types implement.
type serviceProvider interface {
	// setupService sets up the service by supplying it with the [Client].
	setupService(client *Client)
}

// setupServices sets up the Service fields.
func (c *Client) setupServices() {
	// I use reflection here mainly just for convenience: whenever new services
	// are added I can just add them as a field to Client and I don't have to do anything else.
	v := reflect.ValueOf(c).Elem()
	for _, fieldName := range serviceFieldNames {
		service := v.FieldByName(fieldName).Addr().Interface().(serviceProvider)
		service.setupService(c)
	}
}

var serviceFieldNames []string = []string{}

// init checks that all Service fields in the Client struct implement serviceProvider
// and saves the service field names for later use.
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
