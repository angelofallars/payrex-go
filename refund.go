package payrex

// Refund resources represent a refunded amount of a paid payment.
//
// API reference: https://docs.payrexhq.com/docs/api/refunds
type Refund struct {
	BaseResource
	Amount      int                `json:"amount"`
	Currency    Currency           `json:"currency"`
	Status      RefundStatus       `json:"status"`
	Description *string            `json:"description"`
	Reason      RefundReason       `json:"reason"`
	Remarks     *string            `json:"remarks"`
	PaymentID   string             `json:"payment_id"`
	Metadata    *map[string]string `json:"metadata"`
}

// RefundStatus enumerates the valid values for the [Refund].Status field.
type RefundStatus string

const (
	RefundStatusSucceeded RefundStatus = "succeeded"
	RefundStatusFailed    RefundStatus = "failed"
	RefundStatusPending   RefundStatus = "pending"
)

// RefundReason enumerates the valid refund reasons.
type RefundReason string

const (
	RefundReasonFraudulent           RefundReason = "fraudulent"
	RefundReasonRequestedByCustomer  RefundReason = "requested_by_customer"
	RefundReasonProductOutOfStock    RefundReason = "product_out_of_stock"
	RefundReasonServiceNotProvided   RefundReason = "service_not_provided"
	RefundReasonProductWasDamaged    RefundReason = "product_was_damaged"
	RefundReasonServiceMisaligned    RefundReason = "service_misaligned"
	RefundReasonWrongProductReceived RefundReason = "wrong_product_received"
	RefundReasonOthers               RefundReason = "others"
)

// ServiceRefunds provides methods to interact with [Refund] resources,
// available in the [Client].Refunds field.
//
// API reference: https://docs.payrexhq.com/docs/api/refunds
type ServiceRefunds struct{ service[Refund] }

func (s *ServiceRefunds) setup() {
	s.path = prefix("/refunds")
}

// Create creates a refund resource.
//
// Endpoint: POST /refunds
//
// API reference: https://docs.payrexhq.com/docs/api/refunds/create
func (s *ServiceRefunds) Create(params *CreateRefundParams) (*Refund, error) {
	return s.create(params)
}

// Update updates a refund resource by ID.
//
// Endpoint: PUT /refunds/:id
//
// API reference: https://docs.payrexhq.com/docs/api/refunds/update
func (s *ServiceRefunds) Update(id string, params *UpdateRefundParams) (*Refund, error) {
	return s.update(id, params)
}

// CreateRefundParams represents the available [ServiceRefunds.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/refunds/create
type CreateRefundParams struct {
	Amount      int                `form:"amount"`
	Currency    Currency           `form:"currency"`
	Description *string            `form:"description"`
	PaymentID   string             `form:"payment_id"`
	Remarks     *string            `form:"remarks"`
	Reason      RefundReason       `form:"reason"`
	Metadata    *map[string]string `form:"metadata"`
}

// UpdateRefundParams represents the available [ServiceRefunds.Update] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/refunds/update
type UpdateRefundParams struct {
	Metadata *map[string]string `form:"metadata"`
}
