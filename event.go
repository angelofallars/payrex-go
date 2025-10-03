package payrex

import "errors"

// Event represents updates in your PayRex account triggered either by API calls or your actions from the Dashboard.
//
// Service: [ServiceWebhooks] (returned from [ServiceWebhooks.ParseEvent])
//
// API reference: https://docs.payrexhq.com/docs/api/events
type Event struct {
	Resource
	Type EventType `json:"type"`
	// The name of the type of resource associated with this event.
	// Corresponds to the first part of the Type.
	ResourceType       EventResourceType
	PendingWebhooks    int            `json:"pending_webhooks"`
	PreviousAttributes map[string]any `json:"previous_attributes"`

	billingStatement *BillingStatement
	checkoutSession  *CheckoutSession
	paymentIntent    *PaymentIntent
	payout           *Payout
	refund           *Refund
}

// BillingStatement returns the billing statement associated with this event,
// if the event type starts with 'billing_statement'.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a billing statement.
func (e *Event) BillingStatement() (*BillingStatement, error) {
	if e.billingStatement == nil {
		return nil, errors.New("billing statement in event not found")
	}
	return e.billingStatement, nil
}

// CheckoutSession returns the checkout session associated with this event,
// if the event type starts with 'checkout_session'.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a checkout session.
func (e *Event) CheckoutSession() (*CheckoutSession, error) {
	if e.checkoutSession == nil {
		return nil, errors.New("checkout session in event not found")
	}
	return e.checkoutSession, nil
}

// PaymentIntent returns the payment intent associated with this event,
// if the event type starts with 'payment_intent'.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a payment intent.
func (e *Event) PaymentIntent() (*PaymentIntent, error) {
	if e.paymentIntent == nil {
		return nil, errors.New("payment intent in event not found")
	}
	return e.paymentIntent, nil
}

// Payout returns the payout associated with this event,
// if the event type starts with 'payout'.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a payout.
func (e *Event) Payout() (*Payout, error) {
	if e.payout == nil {
		return nil, errors.New("payout in event not found")
	}
	return e.payout, nil
}

// Refund returns the refund associated with this event,
// if the event type starts with 'refund'.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a refund.
func (e *Event) Refund() (*Refund, error) {
	if e.refund == nil {
		return nil, errors.New("refund in event not found")
	}
	return e.refund, nil
}

// MustBillingStatement returns the billing statement associated with this event,
// or panics if there is none.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a billing statement.
func (e *Event) MustBillingStatement() *BillingStatement {
	billingStatement, err := e.BillingStatement()
	if err != nil {
		panic(err)
	}
	return billingStatement
}

// MustCheckoutSession returns the checkout session associated with this event,
// or panics if there is none.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a checkout session.
func (e *Event) MustCheckoutSession() *CheckoutSession {
	checkoutSession, err := e.CheckoutSession()
	if err != nil {
		panic(err)
	}
	return checkoutSession
}

// MustPaymentIntent returns the payment intent associated with this event,
// or panics if there is none.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a payment intent.
func (e *Event) MustPaymentIntent() *PaymentIntent {
	paymentIntent, err := e.PaymentIntent()
	if err != nil {
		panic(err)
	}
	return paymentIntent
}

// MustPayout returns the payout associated with this event,
// or panics if there is none.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a payout.
func (e *Event) MustPayout() *Payout {
	payout, err := e.Payout()
	if err != nil {
		panic(err)
	}
	return payout
}

// MustRefund returns the refund associated with this event,
// or panics if there is none.
//
// It's recommended to check the [Event].ResourceName first to see if the [EventType]
// is of a refund.
func (e *Event) MustRefund() *Refund {
	refund, err := e.Refund()
	if err != nil {
		panic(err)
	}
	return refund
}

// EventType enumerates the event types of an [Event]
// that a [Webhook] can listen to.
type EventType string

const (
	EventTypeBillingStatementCreated             EventType = "billing_statement.created"
	EventTypeBillingStatementUpdated             EventType = "billing_statement.updated"
	EventTypeBillingStatementDeleted             EventType = "billing_statement.deleted"
	EventTypeBillingStatementFinalized           EventType = "billing_statement.finalized"
	EventTypeBillingStatementSent                EventType = "billing_statement.sent"
	EventTypeBillingStatementMarkedUncollectible EventType = "billing_statement.marked_uncollectible"
	EventTypeBillingStatementVoided              EventType = "billing_statement.voided"
	EventTypeBillingStatementPaid                EventType = "billing_statement.paid"
	EventTypeBillingStatementWillBeDue           EventType = "billing_statement.will_be_due"
	EventTypeBillingStatementOverdue             EventType = "billing_statement.overdue"
	EventTypeBillingStatementLineItemCreated     EventType = "billing_statement_line_item.created"
	EventTypeBillingStatementLineItemUpdated     EventType = "billing_statement_line_item.updated"
	EventTypeBillingStatementLineItemDeleted     EventType = "billing_statement_line_item.deleted"
	EventTypeCheckoutSessionExpired              EventType = "checkout_session.expired"
	EventTypePaymentIntentAwaitingCapture        EventType = "payment_intent.awaiting_capture"
	EventTypePaymentIntentSucceeded              EventType = "payment_intent.succeeded"
	EventTypePayoutDeposited                     EventType = "payout.deposited"
	EventTypeRefundCreated                       EventType = "refund.created"
	EventTypeRefundUpdated                       EventType = "refund.updated"
)

// EventResourceType enumerates the types of resources that an [Event]
// can have.
//
// When an Event.ResourceType field matches one of the values here, it's safe to call the
// corresponding Event.Must<Resource>() method.
//
// For example, if the Event.ResourceType is [EventResourceTypeBillingStatement], you can
// call [Event.MustBillingStatement] without issues.
type EventResourceType string

const (
	EventResourceTypeBillingStatement EventResourceType = "billing_statement"
	EventResourceTypeCheckoutSession  EventResourceType = "checkout_session"
	EventResourceTypePaymentIntent    EventResourceType = "payment_intent"
	EventResourceTypePayout           EventResourceType = "payout"
	EventResourceTypeRefund           EventResourceType = "refund"
)
