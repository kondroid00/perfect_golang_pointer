package main

import (
	"fmt"
)

type A struct {
	i int
}

func test(a2 *A) {
	fmt.Printf("a2 value: %p\n", a2)
	fmt.Printf("a2 address: %p\n", &a2)
}

func main() {
	a1 := new(A)
	fmt.Printf("a1 value: %p\n", a1)
	fmt.Printf("a1 address: %p\n", &a1)
	test(a1)
}
