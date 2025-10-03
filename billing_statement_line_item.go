package payrex

// BillingStatementLineItem is a line item of a [BillingStatement] that pertains
// to a business's products or services.
//
// Service: [ServiceBillingStatementLineItems]
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statement_line_items
type BillingStatementLineItem struct {
	BaseResource
	SecretKey          string `json:"secret_key"`
	BillingStatementID string `json:"billing_statement_id"`
	Description        string `json:"description"`
	UnitPrice          int    `json:"unit_price"`
	Quantity           int    `json:"quantity"`
}

// ServiceBillingStatementLineItems provides methods to interact with [BillingStatementLineItem] resources,
// available in the [Client].BillingStatementLineItems field.
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statement_line_items
type ServiceBillingStatementLineItems struct {
	service[BillingStatementLineItem]
}

func (s *ServiceBillingStatementLineItems) setup() {
	s.path = prefix("/billing_statement_line_items")
}

// Create creates a billing statement line item resource.
//
// Endpoint: POST /billing_statement_line_items
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statement_line_items/create
func (s *ServiceBillingStatementLineItems) Create(params *BillingStatementLineItemCreateParams) (*BillingStatementLineItem, error) {
	return s.create(params)
}

// Update updates a billing statement line item resource by ID.
//
// Endpoint: PUT /billing_statement_line_items/:id
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statement_line_items/update
func (s *ServiceBillingStatementLineItems) Update(id string, params *BillingStatementLineItemUpdateParams) (*BillingStatementLineItem, error) {
	return s.update(id, params)
}

// Delete deletes a billing statement line item resource by ID.
//
// Endpoint: DELETE /billing_statement_line_items/:id
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statement_line_items/delete
func (s *ServiceBillingStatementLineItems) Delete(id string) (*DeletedResource, error) {
	return s.delete(id)
}

// BillingStatementLineItemCreateParams represents the available [ServiceBillingStatementLineItems.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statement_line_items/create
type BillingStatementLineItemCreateParams struct {
	BillingStatementID string `form:"billing_statement_id"`
	Description        string `form:"description"`
	UnitPrice          int    `form:"unit_price"`
	Quantity           int    `form:"quantity"`
}

// BillingStatementLineItemUpdateParams represents the available [ServiceBillingStatementLineItems.Update] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statement_line_items/update
type BillingStatementLineItemUpdateParams struct {
	Description *string `form:"description"`
	UnitPrice   *int    `form:"unit_price"`
	Quantity    *int    `form:"quantity"`
}
