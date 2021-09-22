package main

import "fmt"

func plus(x, y int) int {
	return x + y
}
func partialPlus(x int) func(int) int {
	return func(y int) int {
		return plus(x, y)
	}
}
func main() {
	plus_one := partialPlus(1)
	fmt.Println(plus_one(5)) //prints 6

	fmt.Println(plus_one(6)) //prints 6

	fmt.Println(plus_one(7)) //prints 6

}
