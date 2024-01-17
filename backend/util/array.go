package util

import (
	"reflect"
)

func InArray(arr interface{}, element interface{}) bool {
	if reflect.TypeOf(arr).Kind() == reflect.Slice {
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			v := s.Index(i)
			e := reflect.ValueOf(element)
			if v.Equal(e) {
				return true
			}
		}
	}
	return false
}

func RemoveElement(arr interface{}, element interface{}) []interface{} {
	var newArr []interface{}
	if reflect.TypeOf(arr).Kind() == reflect.Slice {
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			v := s.Index(i)
			e := reflect.ValueOf(element)
			if !v.Equal(e) {
				newArr = append(newArr, v.Interface())
			}
		}
	}
	return newArr
}
