package main

import "reflect"

type reducef func(interface{}, interface{}) interface{}

//Reduce(slice, starting value, func)
func Reduce(in interface{}, memo interface{}, fn reducef) interface{} {
	val := reflect.ValueOf(in)

	for i := 0; i < val.Len(); i++ {
		memo = fn(val.Index(i).Interface(), memo)
	}

	return memo
}
