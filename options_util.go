package payrex

// Utility functions for library users to work with Options structs conveniently.

// NotNil returns a pointer to the given value of type T.
//
// Useful for passing in string/int literals to nullable *string/*int fields in Options structs.
//
// In Go 1.26 this will be unnecessary, as the new() function can now accept expressions.
func NotNil[T any](v T) *T {
	return &v
}

// Slice returns a slice containing all the passed in values.
//
// Useful for assigning a slice literal to an Options struct's slice field with type inference,
// instead of manually declaring the slice's element type.
func Slice[T any](values ...T) []T {
	return values
}

// NotNilSlice returns a slice pointer containing all the passed in values.
//
// Useful for assigning a slice literal to an Options struct's nullable slice field with type inference,
// instead of manually declaring the slice's element type.
func NotNilSlice[T any](values ...T) *[]T {
	return &values
}
