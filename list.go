package payrex

// List contains several resources returned together by an API response.
type List[T any] struct {
	// The resources returned by an API response.
	Data []T `json:"data"`
	// Whether there are more resources available.
	HasMore bool `json:"has_more"`
}
