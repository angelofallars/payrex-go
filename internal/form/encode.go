// Package form provides functionality to parse a struct into a URL-encoded string.
package form

import (
	"fmt"
	"net/url"
	"reflect"
)

// Encode returns the URL-encoded form of a struct value
// using `form:"<value>"` tags.
func Encode(params any) string {
	return parseIntoForm(url.Values{}, "", reflect.ValueOf(params)).Encode()
}

func parseIntoForm(formValues url.Values, key string, value reflect.Value) url.Values {
	switch value.Kind() {
	case reflect.Struct:
		valueType := value.Type()
		for i := range value.NumField() {
			field := valueType.Field(i)

			tagKey := field.Tag.Get("form")
			if tagKey == "" {
				panic(fmt.Sprintf("'form' tag on struct field '%s.%s' not set",
					valueType.Name(), field.Name))
			}

			fieldKey := tagKey
			if key != "" {
				fieldKey = fmt.Sprintf("%s[%s]", key, tagKey)
			}
			fieldValue := value.Field(i)

			parseIntoForm(formValues, fieldKey, fieldValue)
		}
	case reflect.Slice:
		for i := range value.Len() {
			elem := value.Index(i)
			elemKey := fmt.Sprintf("%s[]", key)

			parseIntoForm(formValues, elemKey, elem)
		}
	case reflect.Map:
		iter := reflect.ValueOf(value).MapRange()
		for iter.Next() {
			k, v := iter.Key(), iter.Value()

			mapKey := k.String()
			if key != "" {
				mapKey = fmt.Sprintf("%s[%s]", key, k.String())
			}

			parseIntoForm(formValues, mapKey, v)
		}
	case reflect.Pointer:
		// Null values will not be added to the form
		if value.IsZero() {
			return formValues
		}

		parseIntoForm(formValues, key, value.Elem())
	default:
		formValues.Add(key, fmt.Sprintf("%v", value))
	}

	return formValues
}
