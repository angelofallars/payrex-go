package payrex

// serviceProvider is the interface that all Service types implement.
type serviceProvider interface {
	// setupClient sets up the service by supplying it with the [Client].
	setupClient(*Client)
	// setup sets up the service's service-specific data.
	setup()
}

// service is the base type that all Service types embed.
type service[T any] struct {
	client *Client
	path   pathPrefix
}

func (s *service[T]) setupClient(client *Client) {
	s.client = client
}

// Service types embedding service[T] will implement setup()
// on their own to set up any service-specific configuration.
