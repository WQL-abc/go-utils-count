// Package count - counting
package count

import (
	"fmt"
	"reflect"
	"strings"
)

// Do - counting
// supports type: string, slice, array
func Do(something, target interface{}) (int, error) {
	if something == nil {
		return 0, fmt.Errorf("nil does not allow")
	}
	switch reflect.TypeOf(something).Kind() {
	case reflect.String:
		if typ := reflect.TypeOf(target).Kind(); typ != reflect.String {
			return 0, fmt.Errorf("unsupported type(string, %s)", typ.String())
		}
		return strings.Count(something.(string), target.(string)), nil
	case reflect.Slice, reflect.Array:
		targetValue := reflect.ValueOf(target)
		if targetValue.IsValid() && !targetValue.Type().Comparable() {
			return 0, fmt.Errorf("can not comparing uncomparable type")
		}
		somethingValue := reflect.ValueOf(something)
		var cnt int
		for i := 0; i < somethingValue.Len(); i++ {
			indexValue := somethingValue.Index(i)
			if !indexValue.Type().Comparable() {
				continue
			}
			if indexValue.Interface() == target {
				cnt++
			}
		}
		return cnt, nil
	}
	return 0, fmt.Errorf("unsupported type")
}
