package payrex

// TODO: add doc comments for resources when PayRex adds official docs for CustomerSession

// CustomerSession represents a customer session.
//
// Service: [CustomerSession]
type CustomerSession struct {
	Resource
	CustomerID   string                     `json:"customer_id"`
	ClientSecret string                     `json:"client_secret"`
	Components   []CustomerSessionComponent `json:"components"`
	Expired      bool                       `json:"expired"`
	ExpiredAt    int                        `json:"expired_at"`
}

type CustomerSessionComponent struct {
	Component string                        `json:"component"`
	Feature   string                        `json:"feature"`
	Value     CustomerSessionComponentValue `json:"value"`
}

type CustomerSessionComponentValue string

const (
	CustomerSessionComponentValueEnabled  CustomerSessionComponentValue = "enabled"
	CustomerSessionComponentValueDisabled CustomerSessionComponentValue = "disabled"
)

// ServiceCustomerSessions is used to interact with [CustomerSession] resources,
// using the /customer_sessions APIs.
//
// API reference: https://docs.payrexhq.com/docs/api/customer_sessions
type ServiceCustomerSessions struct{ service[CustomerSession] }

func (s *ServiceCustomerSessions) setup() {
	s.path = prefix("/customer_sessions")
}

// Create creates a customer session resource.
//
// Endpoint: POST /customer_sessions
//
// API reference: https://docs.payrexhq.com/docs/api/customer_sessions/create
func (s *ServiceCustomerSessions) Create(params *CustomerSessionCreateParams) (*CustomerSession, error) {
	return s.create(params)
}

// Retrieve retrieves a customer session resource by ID.
//
// Endpoint: GET /customer_sessions/:id
//
// API reference: https://docs.payrexhq.com/docs/api/customer_sessions/retrieve
func (s *ServiceCustomerSessions) Retrieve(id string) (*CustomerSession, error) {
	return s.retrieve(id)
}

// CustomerSessionCreateParams represents the available [ServiceCustomerSessions.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/customer_sessions/create
type CustomerSessionCreateParams struct {
	CustomerID string `form:"customer_id"`
}
