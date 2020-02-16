package main

import "fmt"

type A struct {
	i int
}

func main() {
	var a1 A //‥‥‥①
	a2 := a1 //‥‥‥②
	fmt.Printf("a1: %p\n", &a1)
	fmt.Printf("a2: %p\n", &a2)
}
