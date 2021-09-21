package main

import "fmt"

// New returns a function Count.
// Count prints the number of times it has been invoked.
func New() (Count func()) {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func main() {
	f1, f2 := New(), New()
	f1() // 1
	f2() // 1 (different n)
	f1() // 2
	f2() // 2
	f1() // 2
	f2() // 2
        f1()

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
