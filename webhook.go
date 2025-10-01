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
	Events      []Event       `json:"events"`
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
func (s *ServiceWebhooks) Create(options *CreateWebhookOptions) (*Webhook, error) {
	return s.create(options)
}

// Retrieve retrieves a webhook resource by ID.
//
// Endpoint: GET /webhooks/:id
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/retrieve
func (s *ServiceWebhooks) Retrieve(id string) (*Webhook, error) {
	return s.retrieve(id)
}

// List lists webhooks. The 'options' parameter can be nil.
//
// Endpoint: GET /webhooks
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/list
func (s *ServiceWebhooks) List(options *ListWebhooksOptions) (*Listing[Webhook], error) {
	return s.list(options)
}

// Update updates a webhook resource by ID.
//
// Endpoint: PUT /webhooks/:id
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/update
func (s *ServiceWebhooks) Update(id string, options *UpdateWebhookOptions) (*Webhook, error) {
	return s.update(id, options)
}

// Enable enables a webhook by ID.
//
// Endpoint: POST /webhooks/:id/enable
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/enable
func (s *ServiceWebhooks) Enable(id string) (*Webhook, error) {
	return s.post(
		s.path.make(id, "enable"),
		nil,
	)
}

// Disable disables a webhook resource by ID.
//
// Endpoint: POST /webhooks/:id/disable
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/disable
func (s *ServiceWebhooks) Disable(id string) (*Webhook, error) {
	return s.post(
		s.path.make(id, "disable"),
		nil,
	)
}

// Delete deletes a webhook resource by ID.
//
// Endpoint: DELETE /webhooks/:id
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/delete
func (s *ServiceWebhooks) Delete(id string) (*DeletedResource, error) {
	return s.delete(id)
}

// CreateWebhookOptions contains options for the [ServiceWebhooks.Create] method.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/create
type CreateWebhookOptions struct {
	URL         string  `query:"url"`
	Description *string `query:"description"`
	Events      []Event `query:"events"`
}

// UpdateWebhookOptions contains options for the [ServiceWebhooks.Update] method.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/update
type UpdateWebhookOptions struct {
	URL         *string  `query:"url"`
	Description *string  `query:"description"`
	Events      *[]Event `query:"events"`
}

// ListWebhooksOptions contains options for the [ServiceWebhooks.List] method.
//
// API reference: https://docs.payrexhq.com/docs/api/webhooks/list
type ListWebhooksOptions struct {
	Limit       *int    `query:"int"`
	Before      *string `query:"before"`
	After       *string `query:"after"`
	URL         *string `query:"url"`
	Description *string `query:"description"`
}
