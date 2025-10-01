package payrex

import "net/http"

// Contains helper methods for common endpoint patterns.

func (s *service[T]) create(options any) (*T, error) {
	if options == nil {
		return nil, ErrNilOption
	}

	return s.post(
		s.path.make(),
		options,
	)
}

func (s *service[T]) retrieve(id string) (*T, error) {
	return request[T](s.client,
		http.MethodGet,
		s.path.make(id),
		nil,
	)
}

func (s *service[T]) update(id string, options any) (*T, error) {
	if options == nil {
		return nil, ErrNilOption
	}

	return s.put(
		s.path.make(id),
		options,
	)
}

func (s *service[T]) delete(id string) (*DeletedResource, error) {
	return request[DeletedResource](s.client,
		http.MethodDelete,
		s.path.make(id),
		nil,
	)
}

func (s *service[T]) list(options any) (*Listing[T], error) {
	return request[Listing[T]](s.client,
		http.MethodGet,
		s.path.make(),
		options,
	)
}

func (s *service[T]) post(path urlPath, options any) (*T, error) {
	return request[T](s.client,
		http.MethodPost,
		path,
		options,
	)
}

func (s *service[T]) postID(id string, path urlPath, options any) (*T, error) {
	return s.post(s.path.make(id, string(path)), options)
}

func (s *service[T]) put(path urlPath, options any) (*T, error) {
	return request[T](s.client,
		http.MethodPut,
		path,
		options,
	)
}
