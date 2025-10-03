package payrex

import (
	"encoding/json"
	"fmt"
)

// Webhook is used to notify your application about events in your PayRex account.
//
// Service: [ServiceWebhooks]
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks
type Webhook struct {
	Resource
	SecretKey   string        `json:"secret_key"`
	Status      WebhookStatus `json:"status"`
	Description *string       `json:"description"`
	URL         string        `json:"url"`
	Events      []EventType   `json:"events"`
}

// WebhookStatus enumerates the valid values for the [Webhook].Status field.
type WebhookStatus string

const (
	WebhookStatusEnabled  WebhookStatus = "enabled"
	WebhookStatusDisabled WebhookStatus = "disabled"
)

// ServiceWebhooks provides methods to interact with [Webhook] resources,
// available in the [Client].Webhooks field.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks
type ServiceWebhooks struct{ service[Webhook] }

func (s *ServiceWebhooks) setup() {
	s.path = prefix("/webhooks")
}

// Create creates a webhook resource.
//
// Endpoint: POST /webhooks
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/create
func (s *ServiceWebhooks) Create(params *WebhookCreateParams) (*Webhook, error) {
	return s.create(params)
}

// Retrieve retrieves a webhook resource by ID.
//
// Endpoint: GET /webhooks/:id
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/retrieve
func (s *ServiceWebhooks) Retrieve(id string) (*Webhook, error) {
	return s.retrieve(id)
}

// List lists webhooks. The 'params' parameter can be nil.
//
// Endpoint: GET /webhooks
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/list
func (s *ServiceWebhooks) List(params *WebhookListParams) (*List[Webhook], error) {
	return s.list(params)
}

// Update updates a webhook resource by ID.
//
// Endpoint: PUT /webhooks/:id
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/update
func (s *ServiceWebhooks) Update(id string, params *WebhookUpdateParams) (*Webhook, error) {
	return s.update(id, params)
}

// Enable enables a webhook by ID.
//
// Endpoint: POST /webhooks/:id/enable
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/enable
func (s *ServiceWebhooks) Enable(id string) (*Webhook, error) {
	return s.postID(id, "enable", nil)
}

// Disable disables a webhook resource by ID.
//
// Endpoint: POST /webhooks/:id/disable
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/disable
func (s *ServiceWebhooks) Disable(id string) (*Webhook, error) {
	return s.postID(id, "disable", nil)
}

// Delete deletes a webhook resource by ID.
//
// Endpoint: DELETE /webhooks/:id
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/delete
func (s *ServiceWebhooks) Delete(id string) (*DeletedResource, error) {
	return s.delete(id)
}

// ParseEvent parses an event from a PayRex webhook.
//
// Reference: https://docs.payrexhq.com/docs/guide/developer_handbook/webhooks
func (s *ServiceWebhooks) ParseEvent(payload []byte, signatureHeader, webhookSecretKey string) (*Event, error) {
	// TODO: authenticate event with signatureHeader and webhookSecretKey

	// eventWithResourceName is used to parse the resource name of an [Event].
	type eventWithResourceName struct {
		Data struct {
			Resource string `json:"resource"`
		} `json:"data"`
	}

	// eventWithResource is used to parse the resource type of an [Event].
	type eventWithResource[T any] struct {
		Data T `json:"data"`
	}

	var event Event
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, fmt.Errorf("could not decode event: %w", err)
	}

	var resourceNameContainer eventWithResourceName
	if err := json.Unmarshal(payload, &resourceNameContainer); err != nil {
		return nil, fmt.Errorf("could not decode event resource: %w", err)
	}

	resourceName := resourceNameContainer.Data.Resource

	switch resourceName {

	case "billing_statement":
		var resourceContainer eventWithResource[BillingStatement]
		if err := json.Unmarshal(payload, &resourceContainer); err != nil {
			return nil, fmt.Errorf("could not decode billing statement: %w", err)
		}

		event.billingStatement = &resourceContainer.Data
		event.ResourceType = EventResourceTypeBillingStatement

	case "checkout_session":
		var resourceContainer eventWithResource[CheckoutSession]
		if err := json.Unmarshal(payload, &resourceContainer); err != nil {
			return nil, fmt.Errorf("could not decode checkout session: %w", err)
		}

		event.checkoutSession = &resourceContainer.Data
		event.ResourceType = EventResourceTypeCheckoutSession

	case "payment_intent":
		var resourceContainer eventWithResource[PaymentIntent]
		if err := json.Unmarshal(payload, &resourceContainer); err != nil {
			return nil, fmt.Errorf("could not decode payment intent: %w", err)
		}

		event.paymentIntent = &resourceContainer.Data
		event.ResourceType = EventResourceTypePaymentIntent

	case "payout":
		var resourceContainer eventWithResource[Payout]
		if err := json.Unmarshal(payload, &resourceContainer); err != nil {
			return nil, fmt.Errorf("could not decode payout: %w", err)
		}

		event.payout = &resourceContainer.Data
		event.ResourceType = EventResourceTypePayout

	case "refund":
		var resourceContainer eventWithResource[Refund]
		if err := json.Unmarshal(payload, &resourceContainer); err != nil {
			return nil, fmt.Errorf("could not decode refund: %w", err)
		}

		event.refund = &resourceContainer.Data
		event.ResourceType = EventResourceTypeRefund

	default:
		return nil, fmt.Errorf("unrecognized event resource: '%s'", resourceName)
	}

	return &event, nil
}

// WebhookCreateParams represents the available [ServiceWebhooks.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/create
type WebhookCreateParams struct {
	URL         string      `form:"url"`
	Description *string     `form:"description"`
	Events      []EventType `form:"events"`
}

// WebhookUpdateParams represents the available [ServiceWebhooks.Update] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/update
type WebhookUpdateParams struct {
	URL         *string      `form:"url"`
	Description *string      `form:"description"`
	Events      *[]EventType `form:"events"`
}

// WebhookListParams represents the available [ServiceWebhooks.List] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/list
type WebhookListParams struct {
	Limit       *int    `form:"int"`
	Before      *string `form:"before"`
	After       *string `form:"after"`
	URL         *string `form:"url"`
	Description *string `form:"description"`
}
