package payrex

// baseService is the base type that all Service types embed.
type baseService struct {
	client *Client
}

func (s *baseService) setupService(client *Client) {
	s.client = client
}

// serviceProvider is the interface that all Service types implement.
type serviceProvider interface {
	setupService(client *Client)
}
