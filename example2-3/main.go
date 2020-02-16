package main

import "fmt"

type A struct {
	i int
}

func main() {
	a := A{i: 1}
	b := a    //‥‥‥①
	b.i = 100 //‥‥‥②
	fmt.Printf("a.i: %d\n", a.i)
	fmt.Printf("b.i: %d\n", b.i)
}
