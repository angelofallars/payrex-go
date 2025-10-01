package payrex

// DeletedResource represents a resource of any type that has been deleted.
type DeletedResource struct {
	// Unique identifier for the resource.
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}
