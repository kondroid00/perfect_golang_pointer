package main

import "fmt"

type A struct {
	i int
}

func main() {
	list := []A{{i: 1}, {i: 2}, {i: 3}}

	pList := make([]*A, 0, len(list))
	for _, v := range list {
		// ポインタ型に変換してSliceに突っ込む
		pList = append(pList, &v)
	}

	for i := range pList {
		fmt.Printf("%p\n", pList[i])
	}
}
