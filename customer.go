package payrex

// Customer represents the customer of your business.
// A customer could be a person or a company.
// Use this resource to track payments that belong to the same customer.
//
// Service: [ServiceCustomers]
//
// API reference: https://docs.payrexhq.com/docs/api/customers
type Customer struct {
	Resource
	BillingStatementPrefix             string             `json:"billing_statement_prefix"`
	Currency                           Currency           `json:"currency"`
	Email                              string             `json:"email"`
	Name                               string             `json:"name"`
	Metadata                           *map[string]string `json:"metadata"`
	NextBillingStatementSequenceNumber string             `json:"next_billing_statement_sequence_number"`
}

// ServiceCustomers provides methods to interact with [Customer] resources,
// available in the [Client].Customers field.
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
func (s *ServiceCustomers) Create(params *CustomerCreateParams) (*Customer, error) {
	return s.create(params)
}

// Retrieve retrieves a customer resource by ID.
//
// Endpoint: GET /customers/:id
//
// API reference: https://docs.payrexhq.com/docs/api/customers/retrieve
func (s *ServiceCustomers) Retrieve(id string) (*Customer, error) {
	return s.retrieve(id)
}

// List lists customers. The 'params' parameter can be nil.
//
// Endpoint: GET /customers
//
// API reference: https://docs.payrexhq.com/docs/api/customers/list
func (s *ServiceCustomers) List(params *CustomerListParams) (*List[Customer], error) {
	return s.list(params)
}

// Update updates a customer resource by ID.
//
// Endpoint: PUT /customers/:id
//
// API reference: https://docs.payrexhq.com/docs/api/customers/update
func (s *ServiceCustomers) Update(id string, params *CustomerUpdateParams) (*Customer, error) {
	return s.update(id, params)
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

// CustomerCreateParams represents the available [ServiceCustomers.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/customers/create
type CustomerCreateParams struct {
	Currency                           Currency           `form:"currency"`
	Name                               string             `form:"name"`
	Email                              string             `form:"email"`
	BillingStatementPrefix             *string            `form:"billing_statement_prefix"`
	NextBillingStatementSequenceNumber *string            `form:"next_billing_statement_sequence_number"`
	Metadata                           *map[string]string `form:"metadata"`
}

// CustomerListParams represents the available [ServiceCustomers.List] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/customers/list
type CustomerListParams struct {
	Limit    *int               `form:"int"`
	Before   *string            `form:"before"`
	After    *string            `form:"after"`
	Email    *string            `form:"email"`
	Name     *string            `form:"name"`
	Metadata *map[string]string `form:"metadata"`
}

// CustomerUpdateParams represents the available [ServiceCustomers.Update] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/customers/update
type CustomerUpdateParams struct {
	Currency                           *Currency          `form:"currency"`
	Name                               *string            `form:"name"`
	Email                              *string            `form:"email"`
	BillingStatementPrefix             *string            `form:"billing_statement_prefix"`
	NextBillingStatementSequenceNumber *string            `form:"next_billing_statement_sequence_number"`
	Metadata                           *map[string]string `form:"metadata"`
}
