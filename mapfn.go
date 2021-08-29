package main

import "reflect"

//function types
type mapf func(interface{}) interface{}

//Map(slice, func)
func Map(in interface{}, fn mapf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, val.Len())

	for i := 0; i < val.Len(); i++ {
		out[i] = fn(val.Index(i).Interface())
	}

	return out
}
