package payrex

import "net/http"

// Contains helper methods for common endpoint patterns.

func (s *service[T]) create(params any) (*T, error) {
	if params == nil {
		return nil, ErrNilParams
	}

	return s.post(s.path.make(), params)
}

func (s *service[T]) retrieve(id string) (*T, error) {
	return request[T](s.client,
		http.MethodGet,
		s.path.make(id),
		nil,
	)
}

func (s *service[T]) update(id string, params any) (*T, error) {
	if params == nil {
		return nil, ErrNilParams
	}

	return s.put(s.path.make(id), params)
}

func (s *service[T]) delete(id string) (*DeletedResource, error) {
	return request[DeletedResource](s.client,
		http.MethodDelete,
		s.path.make(id),
		nil,
	)
}

func (s *service[T]) list(params any) (*Listing[T], error) {
	return request[Listing[T]](s.client,
		http.MethodGet,
		s.path.make(),
		params,
	)
}

func (s *service[T]) post(path urlPath, params any) (*T, error) {
	return request[T](s.client,
		http.MethodPost,
		path,
		params,
	)
}

func (s *service[T]) postID(id string, path urlPath, params any) (*T, error) {
	return s.post(s.path.make(id, string(path)), params)
}

func (s *service[T]) put(path urlPath, params any) (*T, error) {
	return request[T](s.client,
		http.MethodPut,
		path,
		params,
	)
}
