package payrex

// BaseResource defines the common fields for most resources.
type BaseResource struct {
	// Unique identifier for the resource.
	ID string `json:"id"`
	// 'true' if the resource's mode is live or 'false' if the resource is in test mode.
	Livemode bool `json:"livemode"`
	// The time the resource was created, measured in seconds since the Unix epoch.
	CreatedAt int `json:"created_at"`
	// The time the resource was updated, measured in seconds since the Unix epoch.
	UpdatedAt int `json:"updated_at"`
}
