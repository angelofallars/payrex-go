package payrex

// CheckoutSession is used to notify your application about events in your PayRex account.
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions
type CheckoutSession struct {
	BaseResource
	URL                      string                    `json:"url"`
	BillingDetailsCollection string                    `json:"billing_details_collection"`
	CustomerReferenceID      *string                   `json:"customer_reference_id"`
	ClientSecret             string                    `json:"client_secret"`
	Status                   CheckoutSessionStatus     `json:"status"`
	Currency                 Currency                  `json:"currency"`
	LineItems                []CheckoutSessionLineItem `json:"line_items"`
	PaymentIntent            *PaymentIntent            `json:"payment_intent"` // TODO: refine type
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

// ServiceCheckoutSessions provides methods to interact with [CheckoutSession] resources,
// available in the [Client].CheckoutSessions field.
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
func (s *ServiceCheckoutSessions) Create(params *CreateCheckoutSessionParams) (*CheckoutSession, error) {
	return s.create(params)
}

// List lists checkout sessions. The 'params' parameter can be nil.
//
// Endpoint: GET /checkout_sessions
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions/list
func (s *ServiceCheckoutSessions) List(params *ListCheckoutSessionsParams) (*Listing[CheckoutSession], error) {
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

// CreateCheckoutSessionParams represents the available [ServiceCheckoutSessions.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/checkout_sessions/create
type CreateCheckoutSessionParams struct {
	URL         string      `form:"url"`
	Description *string     `form:"description"`
	Events      []EventType `form:"events"`
}

// ListCheckoutSessionsParams represents the available [ServiceCheckoutSessions.List] parameters.
type ListCheckoutSessionsParams struct {
	Limit  *int    `form:"int"`
	Before *string `form:"before"`
	After  *string `form:"after"`
}
