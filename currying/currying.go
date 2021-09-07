package main

import "fmt"

func main() {
	in := 10
	m := multiply(4)  // example data
	s := subtract(10) // example data
	// subtract then multiply
	sm := func(i int) int { return m(s(i)) }
	ms := func(i int) int { return s(m(i)) }
	fmt.Printf("%v %v\n", ms(in), sm(in))
}
func multiply(a int) func(int) int {
	return func(i int) int {
		return a * i
	}
}
func subtract(a int) func(int) int {
	return func(i int) int {
		return i - a
	}
}
