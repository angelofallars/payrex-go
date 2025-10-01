package payrex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/angelofallars/payrex-go/internal/query"
)

// request makes a request to the PayRex API with the given payload,
// and returns the JSON response parsed into a value.
func request[T any](client *Client, method method, path urlPath, payload any) (*T, error) {
	reqURL := client.APIBaseURL + string(path)

	var req *http.Request
	var err error

	isPayloadNil := payload == nil || (reflect.ValueOf(payload).Kind() == reflect.Pointer && reflect.ValueOf(payload).IsNil())
	if !isPayloadNil {
		encodedPayload := query.Encode(payload)

		switch method {
		// Put payload in request body
		case methodPOST, methodPUT:
			req, err = http.NewRequest(string(method), reqURL,
				bytes.NewBuffer([]byte(encodedPayload)))
		// Put payload in query parameters
		default:
			req, err = http.NewRequest(string(method), reqURL,
				nil)
			req.URL.RawQuery = encodedPayload
		}
	} else {
		req, err = http.NewRequest(string(method), reqURL,
			nil)
	}

	if err != nil {
		return nil, fmt.Errorf("could not build request: %w", err)
	}

	req.SetBasicAuth(client.apiKey, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
