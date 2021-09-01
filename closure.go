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
}
