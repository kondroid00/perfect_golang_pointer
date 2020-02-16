package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("bool: %d\n", unsafe.Sizeof(false))
	fmt.Printf("byte: %d\n", unsafe.Sizeof(byte(0)))
	fmt.Printf("uint: %d\n", unsafe.Sizeof(uint(0)))
	fmt.Printf("uint8: %d\n", unsafe.Sizeof(uint8(0)))
	fmt.Printf("uint16: %d\n", unsafe.Sizeof(uint16(0)))
	fmt.Printf("uint32: %d\n", unsafe.Sizeof(uint32(0)))
	fmt.Printf("uint64: %d\n", unsafe.Sizeof(uint64(0)))
	fmt.Printf("uintptr: %d\n", unsafe.Sizeof(uintptr(0)))
	fmt.Printf("int: %d\n", unsafe.Sizeof(0))
	fmt.Printf("int8: %d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int16: %d\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("int32: %d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int64: %d\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("float32: %d\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("float64: %d\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("complex64: %d\n", unsafe.Sizeof(complex64(0)))
	fmt.Printf("complex128: %d\n", unsafe.Sizeof(complex128(0)))
	fmt.Printf("string: %d\n", unsafe.Sizeof(""))
	fmt.Printf("rune: %d\n", unsafe.Sizeof('a'))
}
