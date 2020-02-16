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

func getData() *Data {
	return onMemoryData
}

func main() {
	initData() //‥‥‥①

	list := make([]Data, 5)
	for i := range list {
		data := getData() //‥‥‥②
		data.number += i  //‥‥‥③
		list[i] = *data   //‥‥‥④
	}

	fmt.Printf("%v", list)
}
