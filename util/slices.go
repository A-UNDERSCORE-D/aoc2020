package util

import (
	"fmt"
	"reflect"
)

func SliceContains(slice interface{}, value interface{}) bool {
	sType := reflect.TypeOf(slice)
	vType := reflect.TypeOf(value)
	if sType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("SliceContains passed non slice %T", slice))
	}

	if sType.Elem().Kind() != vType.Kind() {
		panic(fmt.Sprintf("SliceContains passed non-matching slice and value types %T, %T", slice, value))
	}

	rSlice := reflect.ValueOf(slice)
	for i := 0; i < rSlice.Len(); i++ {
		v := rSlice.Index(i)
		if reflect.DeepEqual(v.Interface(), value) {
			return true
		}
	}
	return false
}

func StringSliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IntSliceContains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func MakeIntSlice(length, step int) []int {
	num := 0
	out := make([]int, length)
	for i := 0; i < length; i++ {
		out[i] = num
		num += step
	}
	return out
}
