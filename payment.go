package payrex

// Payment represents an individual attempt to move money to your PayRex merchant account balance.
//
// Service: [ServicePayments]
//
// API reference: https://docs.payrexhq.com/docs/api/payments
type Payment struct {
	BaseResource
	Amount          int                `json:"amount"`
	AmountRefunded  int                `json:"amount_refunded"`
	Billing         Billing            `json:"billing"`
	Currency        Currency           `json:"currency"`
	Description     *string            `json:"description"`
	Fee             int                `json:"fee"`
	Metadata        *map[string]string `json:"metadata"`
	NetAmount       int                `json:"net_amount"`
	PaymentIntentID string             `json:"payment_intent_id"`
	Status          PaymentStatus      `json:"payment_status"`
	Customer        *Customer          `json:"customer"`
	PaymentMethod   PaymentMethodType  `json:"payment_method"`
	Refunded        bool               `json:"refunded"`
}

type Billing struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   *string `json:"phone"`
	Address Address `json:"address"`
}

type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"PostalCode"`
	Country    string `json:"country"`
}

type PaymentMethodType struct {
	Type PaymentMethod `json:"type"`
}

// PaymentStatus enumerates the valid values for the [Payment].Status field.
type PaymentStatus string

const (
	PaymentStatusPaid   PaymentStatus = "paid"
	PaymentStatusFailed PaymentStatus = "failed"
)

// ServicePayments provides methods to interact with [Payment] resources, available in the [Client].Payments field.
//
// API reference: https://docs.payrexhq.com/docs/api/payments
type ServicePayments struct{ service[Payment] }

func (s *ServicePayments) setup() {
	s.path = prefix("/payments")
}

// Retrieve retrieves a Payment resource by ID.
//
// Endpoint: GET /payments/:id
//
// API reference: https://docs.payrexhq.com/docs/api/payments/retrieve
func (s *ServicePayments) Retrieve(id string) (*Payment, error) {
	return s.retrieve(id)
}

// Update updates a Payment resource by ID.
//
// Endpoint: PUT /payments/:id
//
// API reference: https://docs.payrexhq.com/docs/api/payments/update
func (s *ServicePayments) Update(id string, options *UpdatePaymentOptions) (*Payment, error) {
	return s.update(id, options)
}

type UpdatePaymentOptions struct {
	Description *string            `query:"description"`
	Metadata    *map[string]string `query:"metadata"`
}
