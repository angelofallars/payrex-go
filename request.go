package payrex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// request makes a request to the PayRex API with the given payload,
// and returns the JSON response parsed into a value.
func request[T any](client *Client, method method, path string, payload any) (*T, error) {
	reqURL := client.APIBaseURL + "/" + path

	var req *http.Request
	var err error

	isPayloadNil := payload == nil || (reflect.ValueOf(payload).Kind() == reflect.Pointer && reflect.ValueOf(payload).IsNil())
	if !isPayloadNil {
		query := buildQuery(payload)

		switch method {
		// Put payload in request body
		case methodPOST, methodPUT:
			req, err = http.NewRequest(string(method), reqURL,
				bytes.NewBuffer([]byte(query)))
		// Put payload in query parameters
		default:
			req, err = http.NewRequest(string(method), reqURL,
				nil)
			req.URL.RawQuery = query
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

func buildQuery(params any) string {
	return buildQueryRecursive(params, "").Encode()
}

func buildQueryRecursive(params any, parentKey string) url.Values {
	queryParams := url.Values{}

	paramsValue := reflect.ValueOf(params)
	if paramsValue.Kind() == reflect.Pointer {
		paramsValue = paramsValue.Elem()
	}

	paramsType := paramsValue.Type()

	for i := range paramsValue.NumField() {
		field := paramsType.Field(i)
		key := field.Tag.Get("query")
		if key == "" {
			panic(fmt.Sprintf("'query' tag on struct field '%s.%s' not set",
				paramsType.Name(), field.Name))
		}

		newKey := key
		if parentKey != "" {
			newKey = fmt.Sprintf("%s[%s]", parentKey, key)
		}

		value := paramsValue.Field(i)
		addQueryValue(queryParams, newKey, value)
	}

	return queryParams
}

func addQueryValue(queryParams url.Values, key string, value reflect.Value) {
	switch value.Kind() {
	case reflect.Struct:
		for k, vs := range buildQueryRecursive(value.Interface(), key) {
			for _, v := range vs {
				queryParams.Add(k, v)
			}
		}
	case reflect.Slice:
		for i := range value.Len() {
			elem := value.Index(i)
			elemKey := fmt.Sprintf("%s[]", key)
			addQueryValue(queryParams, elemKey, elem)
		}
	case reflect.Map:
		for _, k := range value.MapKeys() {
			v := value.MapIndex(k)
			newKey := fmt.Sprintf("%s[%s]", key, k.String())
			addQueryValue(queryParams, newKey, v)
		}
	case reflect.Pointer:
		// Null values will not be added to the query
		if value.IsZero() {
			return
		}

		addQueryValue(queryParams, key, value.Elem())
	default:
		queryParams.Add(key, fmt.Sprintf("%v", value))
	}
}
