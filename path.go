package payrex

import "strings"

// makePath makes a URL path by joining several strings into a single string separated by '/'.
func makePath(paths ...string) string {
	return strings.Join(paths, "/")
}
