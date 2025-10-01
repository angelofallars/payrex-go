package payrex

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
		methodGET,
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
		methodDELETE,
		s.path.make(id),
		nil,
	)
}

func (s *service[T]) list(options any) (*Listing[T], error) {
	return request[Listing[T]](s.client,
		methodGET,
		s.path.make(),
		options,
	)
}

func (s *service[T]) post(path urlPath, payload any) (*T, error) {
	return request[T](s.client,
		methodPOST,
		path,
		payload,
	)
}

func (s *service[T]) put(path urlPath, payload any) (*T, error) {
	return request[T](s.client,
		methodPUT,
		path,
		payload,
	)
}
