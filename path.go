package payrex

import (
	"path"
)

type pathPrefix string

func prefix(path string) pathPrefix {
	if len(path) == 0 || path[0] != '/' {
		panic("empty prefix path")
	}
	if path[0] != '/' {
		panic("path does not start with '/': " + path)
	}

	return pathPrefix(path)
}

type urlPath string

// make returns a URL path with the prefix and given paths joined together.
func (p pathPrefix) make(paths ...string) urlPath {
	if len(paths) == 0 {
		return urlPath(p)
	}

	return urlPath(path.Join(append([]string{string(p)}, paths...)...))
}
