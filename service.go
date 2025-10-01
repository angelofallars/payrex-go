package payrex

// service is the base type that all Service types embed.
type service struct {
	client *Client
}

func (s *service) setupService(client *Client) {
	s.client = client
}

// serviceProvider is the interface that all Service types implement.
type serviceProvider interface {
	setupService(client *Client)
}
