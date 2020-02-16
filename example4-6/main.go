package main

import (
	"fmt"
)

type A struct {
	i int
}

func main() {
	a1 := A{}
	a2 := new(A)
	var a3 A
	var a4 *A
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a3: %T\n", a3)
	fmt.Printf("a4: %T\n", a4)
}
