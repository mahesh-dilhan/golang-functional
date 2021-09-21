package main

import (
	"fmt"
	"reflect"
)

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

func main() {
	a := []int{1, 2, 3, 4}

	c := []int{4, 8, 12, 16}

	//Multiply everything by 2
	b := Map(a, func(val interface{}) interface{} {
		return val.(int) * 2
	})

	//Shoud be [2,4,6,8]
	fmt.Println("MAP:", a, b)

	b = Map(c, func(val interface{}) interface{} {
		return val.(int) / 2
	})

	fmt.Println("MAP:", c, b)

	b = Map(c, func(val interface{}) interface{} {
		return val.(int) + 2
	})

	fmt.Println("MAP:", c, b)

}
