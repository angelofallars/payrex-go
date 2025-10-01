package payrex

// Listing contains several resources returned together by an API response.
type Listing[T any] struct {
	// The resources returned by an API response.
	Values []T `json:"data"`
	// Whether there are more resources available.
	HasMore bool `json:"has_more"`
}
