package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	i int
	f float64
	s string
}

type B struct {
	s string
	a A
}

func main() {
	a := A{}
	fmt.Printf("A: %d\n", unsafe.Sizeof(a))
	fmt.Printf("A pointer: %d\n", unsafe.Sizeof(&a))

	b := B{}
	fmt.Printf("B: %d\n", unsafe.Sizeof(b))
	fmt.Printf("B pointer: %d\n", unsafe.Sizeof(&b))
}
