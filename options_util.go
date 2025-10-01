package payrex

// Utility functions for library users to work with Options structs conveniently.

// Optional returns a pointer to the given value of type T.
//
// Useful for passing in string/int literals to nullable *string/*int fields in Options structs.
//
// In Go 1.26 this will be unnecessary, as the new() function can now accept expressions.
func Optional[T any](v T) *T {
	return &v
}

// Slice returns a slice containing all the passed in values.
//
// Useful if you want to assign a slice literal to a slice field in an Options struct with inferred types,
// instead of manually declaring the slice's element type.
func Slice[T any](values ...T) []T {
	return values
}
