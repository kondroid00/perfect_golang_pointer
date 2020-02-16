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

func main() {
	a := &A{}
	fmt.Printf("a.i: %d\n", unsafe.Offsetof(a.i))
	fmt.Printf("a.f: %d\n", unsafe.Offsetof(a.f))
	fmt.Printf("a.s: %d\n", unsafe.Offsetof(a.s))
}
