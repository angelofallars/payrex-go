package payrex

// PaymentIntent tracks the customer's payment lifecycle, keeping track of
// any failed payment attempts and ensuring the customer is only charged once.
//
// Service: [ServicePaymentIntents]
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents
type PaymentIntent struct {
	BaseResource
	Amount               int                      `json:"amount"`
	AmountReceived       int                      `json:"amount_received"`
	AmountCapturable     int                      `json:"amount_capturable"`
	ClientSecret         string                   `json:"client_secret"`
	Currency             Currency                 `json:"currency"`
	Description          *string                  `json:"description"`
	Metadata             *map[string]string       `json:"metadata"`
	PaymentMethodID      *string                  `json:"payment_method_id"`
	PaymentMethods       []PaymentMethod          `json:"payment_methods"`
	PaymentMethodOptions *PaymentMethodOptions    `json:"payment_method_options"`
	StatementDescriptor  string                   `json:"statement_descriptor"`
	Status               PaymentIntentStatus      `json:"status"`
	NextAction           *PaymentIntentNextAction `json:"next_action"`
	ReturnURL            *string                  `json:"return_url"`
}

// PaymentIntentStatus enumerates the valid values for the [PaymentIntent].Status field.
type PaymentIntentStatus string

const (
	PaymentIntentStatusAwaitingPaymentMethod PaymentIntentStatus = "awaiting_payment_method"
	PaymentIntentStatusAwaitingNextAction    PaymentIntentStatus = "awaiting_next_action"
	PaymentIntentStatusProcessing            PaymentIntentStatus = "processing"
	PaymentIntentStatusSucceeded             PaymentIntentStatus = "succeeded"
)

type PaymentIntentNextAction struct {
	Type        string `json:"type"`
	RedirectURL string `json:"redirect_url"`
}

// ServicePaymentIntents provides methods to interact with [PaymentIntent] resources,
// available in the [Client].PaymentIntents field.
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents
type ServicePaymentIntents struct{ service[PaymentIntent] }

func (s *ServicePaymentIntents) setup() {
	s.path = prefix("/payment_intents")
}

// Cancel cancels a PaymentIntent resource by ID.
//
// A payment intent with a status of canceled means your customer
// cannot proceed with paying the particular payment intent.
//
// You can only cancel a payment intent with status 'awaiting_payment_method'.
//
// Endpoint: POST /payment_intents/:id/cancel
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents/cancel
func (s *ServicePaymentIntents) Cancel(id string) (*PaymentIntent, error) {
	return s.postID(id, "cancel", nil)
}

// Capture captures a PaymentIntent resource by ID.
//
// Endpoint: POST /payment_intents/:id/capture
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents/capture
func (s *ServicePaymentIntents) Capture(id string, params *PaymentIntentCaptureParams) (*PaymentIntent, error) {
	if params == nil {
		return nil, ErrNilParams
	}

	return s.postID(id, "capture", params)
}

// Create creates a PaymentIntent resource.
//
// Endpoint: POST /payment_intents
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents/create
func (s *ServicePaymentIntents) Create(params *PaymentIntentCreateParams) (*PaymentIntent, error) {
	return s.create(params)
}

// Retrieve retrieves a PaymentIntent resource by ID.
//
// Endpoint: GET /payment_intents/:id
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents/retrieve
func (s *ServicePaymentIntents) Retrieve(id string) (*PaymentIntent, error) {
	return s.retrieve(id)
}

// PaymentIntentCaptureParams represents the available [ServicePaymentIntents.Capture] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents/capture
type PaymentIntentCaptureParams struct {
	Amount int `form:"amount"`
}

// PaymentIntentCreateParams represents the available [ServicePaymentIntents.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/payment_intents/create
type PaymentIntentCreateParams struct {
	Amount               int                   `form:"amount"`
	PaymentMethods       []PaymentMethod       `form:"payment_methods"`
	Currency             Currency              `form:"currency"`
	Description          *string               `form:"description"`
	PaymentMethodOptions *PaymentMethodOptions `form:"payment_method_options"`
}
