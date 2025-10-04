package payrex

import "iter"

// CheckoutSession is used to notify your application about events in your PayRex account.
//
// Service: [ServiceCheckoutSessions]
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions
type CheckoutSession struct {
	Resource
	URL                      string                    `json:"url"`
	BillingDetailsCollection string                    `json:"billing_details_collection"`
	CustomerReferenceID      *string                   `json:"customer_reference_id"`
	ClientSecret             string                    `json:"client_secret"`
	Status                   CheckoutSessionStatus     `json:"status"`
	Currency                 Currency                  `json:"currency"`
	LineItems                []CheckoutSessionLineItem `json:"line_items"`
	PaymentIntent            *PaymentIntent            `json:"payment_intent"`
	Metadata                 *map[string]string        `json:"metadata"`
	SuccessURL               string                    `json:"success_url"`
	CancelURL                string                    `json:"cancel_url"`
	PaymentMethods           []PaymentMethod           `json:"payment_methods"`
	Description              *string                   `json:"description"`
	SubmitType               string                    `json:"submit_type"`
	ExpiresAt                int                       `json:"expires_at"`
}

// CheckoutSessionStatus enumerates the valid values for the [CheckoutSession].Status field.
type CheckoutSessionStatus string

const (
	CheckoutSessionStatusActive    CheckoutSessionStatus = "active"
	CheckoutSessionStatusCompleted CheckoutSessionStatus = "completed"
	CheckoutSessionStatusExpired   CheckoutSessionStatus = "expired"
)

type CheckoutSessionLineItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Amount      int     `json:"amount"`
	Quantity    int     `json:"quantity"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
}

// ServiceCheckoutSessions is used to interact with [CheckoutSession] resources,
// using the /checkout_sessions APIs.
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions
type ServiceCheckoutSessions struct{ service[CheckoutSession] }

func (s *ServiceCheckoutSessions) setup() {
	s.path = prefix("/checkout_sessions")
}

// Create creates a checkout session resource.
//
// Endpoint: POST /checkout_sessions
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions/create
func (s *ServiceCheckoutSessions) Create(params *CheckoutSessionCreateParams) (*CheckoutSession, error) {
	return s.create(params)
}

// List lists checkout sessions. The 'params' parameter can be nil.
//
// Endpoint: GET /checkout_sessions
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions/list
func (s *ServiceCheckoutSessions) List(params *ListCheckoutSessionsParams) iter.Seq2[*CheckoutSession, error] {
	return s.list(params)
}

// Retrieve retrieves a checkout session resource by ID.
//
// A [CheckoutSession] can only be retrieved from the server side using a secret API key.
//
// Endpoint: GET /checkout_sessions/:id
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions/retrieve
func (s *ServiceCheckoutSessions) Retrieve(id string) (*CheckoutSession, error) {
	return s.retrieve(id)
}

// Expire marks a checkout_session as expired by ID.
//
// Endpoint: POST /checkout_sessions/:id/enable
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions/enable
func (s *ServiceCheckoutSessions) Expire(id string) (*CheckoutSession, error) {
	return s.postID(id, "expire", nil)
}

// CheckoutSessionCreateParams represents the available [ServiceCheckoutSessions.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions/create
type CheckoutSessionCreateParams struct {
	CustomerReferenceID      *string                         `form:"customer_reference_id"`
	Currency                 Currency                        `form:"currency"`
	LineItems                []CheckoutSessionLineItemParams `form:"line_items"`
	Metadata                 *map[string]string              `form:"metadata"`
	SuccessURL               string                          `form:"success_url"`
	CancelURL                string                          `form:"cancel_url"`
	ExpiresAt                *int                            `form:"expires_at"`
	PaymentMethods           []PaymentMethod                 `form:"payment_methods"`
	BillingDetailsCollection *string                         `form:"billing_details_collection"`
	Description              *string                         `form:"description"`
	SubmitType               *string                         `form:"submit_type"`
	PaymentMethodOptions     *PaymentMethodOptions           `form:"payment_method_options"`
}

type CheckoutSessionLineItemParams struct {
	Name        string  `form:"name"`
	Amount      int     `form:"amount"`
	Quantity    int     `form:"quantity"`
	Description *string `form:"description"`
	Image       *string `form:"image"`
}

// TODO: add API reference for ListCheckoutSessionsParams when docs for `GET /checkout_sessions` is added

// ListCheckoutSessionsParams represents the available [ServiceCheckoutSessions.List] parameters.
type ListCheckoutSessionsParams struct {
	Limit  *int    `form:"int"`
	Before *string `form:"before"`
	After  *string `form:"after"`
}
