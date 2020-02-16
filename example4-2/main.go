package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("array1: %d\n", unsafe.Sizeof([1]int{1}))
	fmt.Printf("array2: %d\n", unsafe.Sizeof([3]int{1, 1, 1}))
	fmt.Printf("array3: %d\n", unsafe.Sizeof([3]int8{1, 1, 1}))

	fmt.Printf("slice1: %d\n", unsafe.Sizeof([]int{1}))
	fmt.Printf("slice2: %d\n", unsafe.Sizeof([]int{1, 1, 1}))
	fmt.Printf("slice3: %d\n", unsafe.Sizeof([]int8{1, 1, 1}))

	fmt.Printf("map1: %d\n", unsafe.Sizeof(
		map[string]int{"a": 1}))
	fmt.Printf("map2: %d\n", unsafe.Sizeof(
		map[string]int{"a": 1, "b": 2, "c": 3}))
	fmt.Printf("map3: %d\n", unsafe.Sizeof(
		map[string]int8{"a": 1, "b": 2, "c": 3}))

	fmt.Printf("chan1: %d\n", unsafe.Sizeof(make(chan int, 1)))
	fmt.Printf("chan2: %d\n", unsafe.Sizeof(make(chan int, 3)))
	fmt.Printf("chan3: %d\n", unsafe.Sizeof(make(chan int8, 3)))
}
