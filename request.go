package payrex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/angelofallars/payrex-go/internal/form"
)

// request makes a request to the PayRex API with the given payload,
// and returns the JSON response parsed into a value.
func request[T any](client *Client, method string, path urlPath, payload any) (*T, error) {
	reqURL := client.apiBaseURL + string(path)

	var req *http.Request
	var err error

	isPayloadNil := payload == nil || (reflect.ValueOf(payload).Kind() == reflect.Pointer && reflect.ValueOf(payload).IsNil())
	if !isPayloadNil {
		encodedPayload := form.Encode(payload)

		switch method {
		// Put payload in request body
		case http.MethodPost, http.MethodPut:
			req, err = http.NewRequest(string(method), reqURL, bytes.NewBuffer([]byte(encodedPayload)))
		// Put payload in query parameters
		default:
			req, err = http.NewRequest(string(method), reqURL, nil)
			req.URL.RawQuery = encodedPayload
		}
	} else {
		req, err = http.NewRequest(string(method), reqURL, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("could not build request: %w", err)
	}

	req.SetBasicAuth(client.apiKey, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	const userAgentName = "payrex-go"
	const userAgentVersion = "0.0.1"
	req.Header.Set("User-Agent", userAgentName+"/"+userAgentVersion)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not do request: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		var errBody Error
		if err := json.NewDecoder(res.Body).Decode(&errBody); err != nil {
			return nil, fmt.Errorf("could not decode error response JSON: %w", err)
		}

		return nil, errBody
	}

	var resource T
	if err := json.NewDecoder(res.Body).Decode(&resource); err != nil {
		return nil, fmt.Errorf("could not decode JSON: %w", err)
	}

	return &resource, nil
}

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

func (s *service[T]) list(params any) (*List[T], error) {
	return request[List[T]](s.client,
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
