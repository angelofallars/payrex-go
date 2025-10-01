package payrex

// Payout resources are created when you are scheduled to receive money from PayRex.
//
// Payouts are made depending on the payout schedule for your PayRex merchant account.
//
// API reference: https://docs.payrexhq.com/docs/api/payouts
type Payout struct {
	BaseResource
	Amount      int               `json:"amount"`
	Destination PayoutDestination `json:"destination"`
	NetAmount   int               `json:"net_amount"`
	Status      PayoutStatus      `json:"status"`
}

// PayoutStatus enumerates the valid values for the [Payout].Status field.
type PayoutStatus string

const (
	PayoutStatusPending    PayoutStatus = "pending"
	PayoutStatusInTransit  PayoutStatus = "in_transit"
	PayoutStatusFailed     PayoutStatus = "failed"
	PayoutStatusSuccessful PayoutStatus = "successful"
)

type PayoutDestination struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
}

// PayoutTransaction represents every line item of a Payout.
//
// Every [PayoutTransaction] belongs to a [Payout] resource.
//
// API reference: https://docs.payrexhq.com/docs/api/payout_transactions
type PayoutTransaction struct {
	// Unique identifier for the resource.
	ID              string                `json:"id"`
	Amount          int                   `json:"amount"`
	NetAmount       int                   `json:"net_amount"`
	TransactionType PayoutTransactionType `json:"transaction_type"`
	TransactionID   string                `json:"transaction_id"`
	// The time the resource was created, measured in seconds since the Unix epoch.
	CreatedAt int `json:"created_at"`
	// The time the resource was updated, measured in seconds since the Unix epoch.
	UpdatedAt int `json:"updated_at"`
}

// PayoutTransactionType enumerates the valid values for the [PayoutTransaction].TransactionType field.
type PayoutTransactionType string

const (
	PayoutTransactionTypePayment    PayoutTransactionType = "payment"
	PayoutTransactionTypeRefund     PayoutTransactionType = "refund"
	PayoutTransactionTypeAdjustment PayoutTransactionType = "adjustment"
)

// ServicePayouts provides methods to interact with [Payout] and [PayoutTransaction] resources, available in the [Client].Payouts field.
//
// API reference: https://docs.payrexhq.com/docs/api/payouts
type ServicePayouts struct{ service }

func (s *ServicePayouts) setup() {
	s.path = prefix("/payouts")
}

// ListTransactions lists payout transactions by [Payout] ID. The 'options' parameter can be nil.
//
// Endpoint: GET /payouts/:id/transactions
//
// API reference: https://docs.payrexhq.com/docs/api/payout_transactions/list
func (s *ServicePayouts) ListTransactions(id string, options *ListPayoutTransactionsOptions) (*Listing[PayoutTransaction], error) {
	return request[Listing[PayoutTransaction]](s.client,
		methodGET,
		s.path.make(id, "transactions"),
		options,
	)
}

// ListPayoutTransactionsOptions contains options for the [ServicePayouts.ListTransactions] method.
//
// API reference: https://docs.payrexhq.com/docs/api/payout_transactions/list
type ListPayoutTransactionsOptions struct {
	Limit  *int    `query:"limit"`
	Before *string `query:"before"`
	After  *string `query:"after"`
}
