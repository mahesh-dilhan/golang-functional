package main

import (
	"fmt"
	"reflect"
)

type reducef func(interface{}, interface{}) interface{}

//Reduce(slice, starting value, func)
func Reduce(in interface{}, memo interface{}, fn reducef) interface{} {
	val := reflect.ValueOf(in)

	for i := 0; i < val.Len(); i++ {
		memo = fn(val.Index(i).Interface(), memo)
	}

	return memo
}

func main() {
	b := []int{1, 2, 3, 4}
	//Summation
	c := Reduce(b, 0, func(val interface{}, memo interface{}) interface{} {
		return memo.(int) + val.(int)
	})

	//Should be 20
	fmt.Println("REDUCE:", b, c)

	c = Reduce(b, 0, func(val interface{}, memo interface{}) interface{} {
		return memo.(int) - val.(int)
	})

	fmt.Println("REDUCE:", b, c)

}
