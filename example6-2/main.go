package main

import "fmt"

type Data struct {
	number int
}

var onMemoryData *Data

func initData() {
	onMemoryData = &Data{
		number: 100,
	}
}

func getData() Data { //‥‥‥⑤
	return *onMemoryData //‥‥‥⑥
}

func main() {
	initData()

	list := make([]Data, 5)
	for i := range list {
		data := getData() //‥‥‥⑦
		data.number += i  //‥‥‥⑧
		list[i] = data
	}

	fmt.Printf("%v", list)
}
