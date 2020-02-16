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
	a1 := A{}                      //‥‥‥①
	a1.f = 2.4                     //‥‥‥②
	fmt.Printf("a1.f: %f\n", a1.f) //‥‥‥③
	fmt.Printf("a1 address: %p\n", &a1)
	fmt.Printf("a1 size: %d\n", unsafe.Sizeof(a1))
}
