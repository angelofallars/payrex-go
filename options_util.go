package payrex

// Optional returns a pointer to the given value of type T.
//
// Useful for passing in string/int literals to nullable *string/*int fields in Options structs.
//
// In Go 1.26 this will be unnecessary, as the new() function can now accept expressions.
func Optional[T any](v T) *T {
	return &v
}
