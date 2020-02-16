package main

import (
	"fmt"
)

func main() {
	var b *int
	fmt.Printf("value: %p\n", b)
	fmt.Printf("address: %p\n", &b)
}
