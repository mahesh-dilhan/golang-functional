package main

import (
	"fmt"
	"reflect"
)

type filterf func(interface{}) bool

//Filter(slice, predicate func)
func Filter(in interface{}, fn filterf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		current := val.Index(i).Interface()

		if fn(current) {
			out = append(out, current)
		}
	}

	return out
}

func main() {
	b := []int{1, 2, 3, 4}

	//Check if the number is divisble by 4
	d := Filter(b, func(val interface{}) bool {
		return val.(int)%4 == 0
	})

	//Should be [4,8]
	fmt.Println("FILTER:", b, d)

}
