package payrex

// Customer represents the customer of your business. A customer could be a person or a company. Use this resource to track payments that belong to the same customer.
//
// Service: [ServiceCustomers]
//
// API reference: https://docs.payrexhq.com/docs/api/customers
type Customer struct {
	BaseResource
	BillingStatementPrefix             string             `json:"billing_statement_prefix"`
	Currency                           Currency           `json:"currency"`
	Email                              string             `json:"email"`
	Name                               string             `json:"name"`
	Metadata                           *map[string]string `json:"metadata"`
	NextBillingStatementSequenceNumber string             `json:"next_billing_statement_sequence_number"`
}

// ServiceCustomers provides methods to interact with [Customer] resources, available in the [Client].Customers field.
//
// API reference: https://docs.payrexhq.com/docs/api/customers
type ServiceCustomers struct{ service[Customer] }

func (s *ServiceCustomers) setup() {
	s.path = prefix("/customers")
}

// Create creates a customer resource.
//
// Endpoint: POST /customers
//
// API reference: https://docs.payrexhq.com/docs/api/customers/create
func (s *ServiceCustomers) Create(options *CreateCustomerOptions) (*Customer, error) {
	return s.create(options)
}

// Retrieve retrieves a customer resource by ID.
//
// Endpoint: GET /customers/:id
//
// API reference: https://docs.payrexhq.com/docs/api/customers/retrieve
func (s *ServiceCustomers) Retrieve(id string) (*Customer, error) {
	return s.retrieve(id)
}

// List lists customers. The 'options' parameter can be nil.
//
// Endpoint: GET /customers
//
// API reference: https://docs.payrexhq.com/docs/api/customers/list
func (s *ServiceCustomers) List(options *ListCustomersOptions) (*Listing[Customer], error) {
	return s.list(options)
}

// Update updates a customer resource by ID.
//
// Endpoint: PUT /customers/:id
//
// API reference: https://docs.payrexhq.com/docs/api/customers/update
func (s *ServiceCustomers) Update(id string, options *UpdateCustomerOptions) (*Customer, error) {
	return s.update(id, options)
}

// Delete deletes a customer resource by ID.
//
// Deleted customers can still be retrieved through [ServiceCustomers.Retrieve] to track their history.
//
// Endpoint: DELETE /customers/:id
//
// API reference: https://docs.payrexhq.com/docs/api/customers/delete
func (s *ServiceCustomers) Delete(id string) (*DeletedResource, error) {
	return s.delete(id)
}

// CreateCustomerOptions contains options for the [ServiceCustomers.Create] method.
//
// API reference: https://docs.payrexhq.com/docs/api/customers/create
type CreateCustomerOptions struct {
	Currency                           Currency           `query:"currency"`
	Name                               string             `query:"name"`
	Email                              string             `query:"email"`
	BillingStatementPrefix             *string            `query:"billing_statement_prefix"`
	NextBillingStatementSequenceNumber *string            `query:"next_billing_statement_sequence_number"`
	Metadata                           *map[string]string `query:"metadata"`
}

// ListCustomersOptions contains options for the [ServiceCustomers.List] method.
//
// API reference: https://docs.payrexhq.com/docs/api/customers/list
type ListCustomersOptions struct {
	Limit    *int               `query:"int"`
	Before   *string            `query:"before"`
	After    *string            `query:"after"`
	Email    *string            `query:"email"`
	Name     *string            `query:"name"`
	Metadata *map[string]string `query:"metadata"`
}

// UpdateCustomerOptions contains options for the [ServiceCustomers.Update] method.
//
// API reference: https://docs.payrexhq.com/docs/api/customers/update
type UpdateCustomerOptions struct {
	Currency                           *Currency          `query:"currency"`
	Name                               *string            `query:"name"`
	Email                              *string            `query:"email"`
	BillingStatementPrefix             *string            `query:"billing_statement_prefix"`
	NextBillingStatementSequenceNumber *string            `query:"next_billing_statement_sequence_number"`
	Metadata                           *map[string]string `query:"metadata"`
}
