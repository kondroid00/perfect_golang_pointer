package main

import "fmt"

type A struct {
	i int
}

func main() {
	list := []A{{i: 1}, {i: 2}, {i: 3}}
	pList := make([]*A, 0, len(list))

	var v A //‥‥‥①

	v = list[0]               //‥‥‥②
	pList = append(pList, &v) //‥‥‥③

	v = list[1]               //‥‥‥④
	pList = append(pList, &v) //‥‥‥⑤

	v = list[2]               //‥‥‥⑥
	pList = append(pList, &v) //‥‥‥⑦

	for _, v := range pList {
		fmt.Println(v.i) //‥‥‥⑧
	}
}
