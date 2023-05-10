package utils

import (
	"errors"
	"reflect"
)

var (
	ErrUnSupportSlice = errors.New("haystack must be slice, array")
)

func FindInSlice(haystack []string, val string) bool {
	for _, v := range haystack {
		if v == val {
			return true
		}
	}

	return false
}

func FindInSliceByInterface(haystack interface{}, val interface{}) (bool, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == val {
				return true, nil
			}
		}

		return false, nil
	}

	return false, ErrUnSupportSlice
}
