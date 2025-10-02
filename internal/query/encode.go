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
	return parseQueryParams(url.Values{}, "", reflect.ValueOf(params)).Encode()
}

func parseQueryParams(queryParams url.Values, key string, value reflect.Value) url.Values {
	switch value.Kind() {
	case reflect.Struct:
		valueType := value.Type()
		for i := range value.NumField() {
			field := valueType.Field(i)

			tagKey := field.Tag.Get("query")
			if tagKey == "" {
				panic(fmt.Sprintf("'query' tag on struct field '%s.%s' not set",
					valueType.Name(), field.Name))
			}

			fieldKey := tagKey
			if key != "" {
				fieldKey = fmt.Sprintf("%s[%s]", key, tagKey)
			}
			fieldValue := value.Field(i)

			parseQueryParams(queryParams, fieldKey, fieldValue)
		}
	case reflect.Slice:
		for i := range value.Len() {
			elem := value.Index(i)
			elemKey := fmt.Sprintf("%s[]", key)

			parseQueryParams(queryParams, elemKey, elem)
		}
	case reflect.Map:
		iter := reflect.ValueOf(value).MapRange()
		for iter.Next() {
			k, v := iter.Key(), iter.Value()

			mapKey := k.String()
			if key != "" {
				mapKey = fmt.Sprintf("%s[%s]", key, k.String())
			}

			parseQueryParams(queryParams, mapKey, v)
		}
	case reflect.Pointer:
		// Null values will not be added to the query
		if value.IsZero() {
			return queryParams
		}

		parseQueryParams(queryParams, key, value.Elem())
	default:
		queryParams.Add(key, fmt.Sprintf("%v", value))
	}

	return queryParams
}
