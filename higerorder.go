package main

import (
	"fmt"
	"strings"
)

func main() {

	f := func(str string) string {
		return strings.ToUpper(str)
	}

	caller("Golang", f)

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

	var add10 = add(10)
	var add20 = add(20)
	var add30 = add(30)

	fmt.Println(add10(5)) // 15
	fmt.Println(add20(5)) // 25
	fmt.Println(add30(5)) // 35

}

func Mapfn(f func(interface{}) interface{}, xs []interface{}) []interface{} {
	ys := make([]interface{}, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func caller(str string, f func(str string) string) {
	fmt.Println(f("hellow world " + str))
}

// this is a higher-order-function that returns a function
func add(x int) func(y int) int {
	// A function is returned here as closure
	// variable x is obtained from the outer scope of this method and memorized in the closure
	return func(y int) int {
		return x + y
	}
}
