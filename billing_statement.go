package payrex

import "iter"

// TODO: update BillingStatement fields as the BillingStatement official API docs become more accurate

// BillingStatement is used to notify your application about events in your PayRex account.
//
// Service: [ServiceBillingStatements]
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements
type BillingStatement struct {
	Resource
	Status    BillingStatementStatus `json:"status"`
	Amount    int                    `json:"amount"`
	Currency  Currency               `json:"currency"`
	LineItems []struct {
		ID                 string `json:"id"`
		BillingStatementID string `json:"billing_statement_id"`
		Description        string `json:"description"`
		UnitPrice          int    `json:"unit_price"`
		Quantity           int    `json:"quantity"`
	} `json:"line_items"`
	PaymentIntent            *PaymentIntent     `json:"payment_intent"`
	BillingDetailsCollection string             `json:"billing_details_collection"`
	CustomerID               string             `json:"customer_id"`
	Description              *string            `json:"description"`
	MerchantName             *string            `json:"billing_statement_merchant_name"`
	MerchantNumber           *string            `json:"billing_statement_merchant_number"`
	URL                      *string            `json:"billing_statement_url"`
	StatementDescriptor      *string            `json:"statement_descriptor"`
	PaymentSettings          PaymentSettings    `json:"payment_settings"`
	Metadata                 *map[string]string `json:"metadata"`
}

// PaymentSettings lists fields that can modify the behavior of the payment processing for a [BillingStatement].
type PaymentSettings struct {
	// The list of [PaymentMethod] values allowed to be processed by the [PaymentIntent] of a [BillingStatement].
	PaymentMethods []PaymentMethod `json:"payment_methods" form:"payment_methods"`
}

// BillingStatementStatus enumerates the valid values for the [BillingStatement].Status field.
type BillingStatementStatus string

const (
	BillingStatementStatusOpen          BillingStatementStatus = "open"
	BillingStatementStatusDraft         BillingStatementStatus = "draft"
	BillingStatementStatusPaid          BillingStatementStatus = "paid"
	BillingStatementStatusVoid          BillingStatementStatus = "void"
	BillingStatementStatusUncollectible BillingStatementStatus = "uncollectible"
)

// ServiceBillingStatements is used to interact with [BillingStatement] resources,
// using the /billing_statements APIs.
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements
type ServiceBillingStatements struct{ service[BillingStatement] }

func (s *ServiceBillingStatements) setup() {
	s.path = prefix("/billing_statements")
}

// Create creates a billing statement resource.
//
// Endpoint: POST /billing_statements
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/create
func (s *ServiceBillingStatements) Create(params *BillingStatementCreateParams) (*BillingStatement, error) {
	return s.create(params)
}

// Retrieve retrieves a billing statement resource by ID.
//
// Endpoint: GET /billing_statements/:id
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/retrieve
func (s *ServiceBillingStatements) Retrieve(id string) (*BillingStatement, error) {
	return s.retrieve(id)
}

// List lists billing statement resources. The 'params' parameter can be nil.
//
// Endpoint: GET /billing_statements
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/list
func (s *ServiceBillingStatements) List(params *BillingStatementListParams) iter.Seq2[*BillingStatement, error] {
	return s.list(params)
}

// Update updates a billing statement resource by ID.
//
// Endpoint: PUT /billing_statements/:id
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/update
func (s *ServiceBillingStatements) Update(id string, params *BillingStatementUpdateParams) (*BillingStatement, error) {
	return s.update(id, params)
}

// Delete deletes a billing statement resource by ID.
//
// Endpoint: DELETE /billing_statements/:id
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/delete
func (s *ServiceBillingStatements) Delete(id string) (*DeletedResource, error) {
	return s.delete(id)
}

// Finalize finalizes a billing statement by ID.
//
// Endpoint: POST /billing_statements/:id/finalize
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/finalize
func (s *ServiceBillingStatements) Finalize(id string) (*BillingStatement, error) {
	return s.postID(id, "finalize", nil)
}

// MarkUncollectible marks a billing statement resource as uncollectible by ID.
//
// Endpoint: POST /billing_statements/:id/disable
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/mark_uncollectible
func (s *ServiceBillingStatements) MarkUncollectible(id string) (*BillingStatement, error) {
	return s.postID(id, "mark_uncollectible", nil)
}

// Send sends a billing statement resource via e-mail by ID.
//
// Endpoint: POST /billing_statements/:id/disable
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/send
func (s *ServiceBillingStatements) Send(id string) (*BillingStatement, error) {
	return s.postID(id, "send", nil)
}

// Void void a billing statement resource by ID.
//
// Endpoint: POST /billing_statements/:id/void
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/void
func (s *ServiceBillingStatements) Void(id string) (*BillingStatement, error) {
	return s.postID(id, "void", nil)
}

// BillingStatementCreateParams represents the available [ServiceBillingStatements.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/create
type BillingStatementCreateParams struct {
	CustomerID               string             `form:"customer_id"`
	Currency                 Currency           `form:"currency"`
	Description              *string            `form:"description"`
	BillingDetailsCollection *string            `form:"billing_details_collection"`
	PaymentSettings          PaymentSettings    `form:"payment_settings"`
	Metadata                 *map[string]string `form:"metadata"`
}

// BillingStatementUpdateParams represents the available [ServiceBillingStatements.Update] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/update
type BillingStatementUpdateParams struct {
	CustomerID               *string            `form:"customer_id"`
	Description              *string            `form:"description"`
	BillingDetailsCollection *string            `form:"billing_details_collection"`
	PaymentSettings          *PaymentSettings   `form:"payment_settings"`
	Metadata                 *map[string]string `form:"metadata"`
}

// BillingStatementListParams represents the available [ServiceBillingStatements.List] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/billing_statements/list
type BillingStatementListParams struct {
	Limit  *int    `form:"int"`
	Before *string `form:"before"`
	After  *string `form:"after"`
}
