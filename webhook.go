package payrex

// Webhook is used to notify your application about events in your PayRex account.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks
type Webhook struct {
	BaseResource
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
func (s *ServiceWebhooks) Create(params *CreateWebhookParams) (*Webhook, error) {
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
func (s *ServiceWebhooks) List(params *ListWebhooksParams) (*Listing[Webhook], error) {
	return s.list(params)
}

// Update updates a webhook resource by ID.
//
// Endpoint: PUT /webhooks/:id
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/update
func (s *ServiceWebhooks) Update(id string, params *UpdateWebhookParams) (*Webhook, error) {
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

// CreateWebhookParams represents the available [ServiceWebhooks.Create] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/create
type CreateWebhookParams struct {
	URL         string      `form:"url"`
	Description *string     `form:"description"`
	Events      []EventType `form:"events"`
}

// UpdateWebhookParams represents the available [ServiceWebhooks.Update] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/update
type UpdateWebhookParams struct {
	URL         *string      `form:"url"`
	Description *string      `form:"description"`
	Events      *[]EventType `form:"events"`
}

// ListWebhooksParams represents the available [ServiceWebhooks.List] parameters.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/list
type ListWebhooksParams struct {
	Limit       *int    `form:"int"`
	Before      *string `form:"before"`
	After       *string `form:"after"`
	URL         *string `form:"url"`
	Description *string `form:"description"`
}
