package utils

import (
	"reflect"
	"strconv"
)

// InterfacesToStrings convert the interfaces slice to the string slice.
func InterfacesToStrings(d []interface{}) []string {
	r := []string{}

	for _, v := range d {
		switch reflect.ValueOf(v).Kind() {
		case reflect.String:
			r = append(r, v.(string))
			break

		case reflect.Int, reflect.Int32:
			r = append(r, strconv.Itoa(v.(int)))
			break

		default:
			// TODO
		}
	}

	return r
}
