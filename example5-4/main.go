package main

import "fmt"

type A struct {
	i int
}

func main() {
	list := []A{{i: 1}, {i: 2}, {i: 3}}

	pList := make([]*A, 0, len(list))
	for i := range list {
		// インデックスで要素を指定してポインタを入れる
		pList = append(pList, &list[i])
	}

	for _, v := range pList {
		fmt.Println(v.i)
	}
}
