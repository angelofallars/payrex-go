package payrex

import (
	"fmt"
	"reflect"
	"strings"
)

var serviceFieldNames []string = []string{}

// setupServices sets up the Service fields.
func (c *Client) setupServices() {
	// I use reflection here mainly just for convenience: whenever new services
	// are added I can just add them as a field to Client and I don't have to do anything else.
	v := reflect.ValueOf(c).Elem()
	for _, fieldName := range serviceFieldNames {
		service := v.FieldByName(fieldName).Addr().Interface().(serviceProvider)

		service.setupClient(c)
		service.setup()
	}
}

// init checks that all Service fields in the [Client] struct implement [serviceProvider],
// and saves the service field names in [serviceFieldNames] for later use.
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
