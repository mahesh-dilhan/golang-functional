package main

import (
	"fmt"
	"strings"
)

func main() {

	f := func(str string) string {
		return strings.ToUpper(str)
	}

	caller(f)

	square := func(x interface{}) interface{} {
		return x.(int) * x.(int)
	}
	nums := []int{1, 2, 3, 4}

	fmt.Println(nums)
	gnums := make([]interface{}, len(nums))
	for i, x := range nums {
		gnums[i] = x
	}
	fmt.Println(gnums)

	gsquared := Mapfn(square, gnums)
	squared := make([]int, len(gsquared))
	for i, x := range gsquared {
		squared[i] = x.(int)
	}

	fmt.Println(squared)
}

func Mapfn(f func(interface{}) interface{}, xs []interface{}) []interface{} {
	ys := make([]interface{}, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func caller(f func(str string) string) {
	fmt.Println(f("hellow world"))
}
