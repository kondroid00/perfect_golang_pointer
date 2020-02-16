package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	i int
	f float64
}

func main() {
	a2 := new(A)                   //‥‥‥①
	a2.f = 2.4                     //‥‥‥②
	fmt.Printf("a2.f: %f\n", a2.f) //‥‥‥③
	fmt.Printf("a2 address: %p\n", &a2)
	fmt.Printf("a2 value: %p\n", a2)
	fmt.Printf("a2 size: %d\n", unsafe.Sizeof(*a2))
}
