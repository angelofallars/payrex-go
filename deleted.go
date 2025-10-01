package payrex

// DeletedResource represents a resource of any type that has been deleted.
type DeletedResource struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}
