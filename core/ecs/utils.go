package ecs

import "reflect"

// intInSlice will return true if the integer value provided is present in the slice provided, false otherwise.
func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// typeInSlice will return true if the reflect.Type provided is present in the slice provided, false otherwise.
func typeInSlice(a reflect.Type, list []reflect.Type) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
