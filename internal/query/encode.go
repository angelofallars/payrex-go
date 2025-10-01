// Package query provides functionality to parse a struct into URL query values.
package query

import (
	"fmt"
	"net/url"
	"reflect"
)

// Encode returns a URL encoded version of a struct value
// using `query:"<value>"` tags.
func Encode(params any) string {
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
