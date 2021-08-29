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
}

func caller(f func(str string) string) {
	fmt.Println(f("hellow world"))
}
