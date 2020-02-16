package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Number1 int  `json:"number1"`
	Number2 *int `json:"number2"`
}

func main() {
	b1 := []byte(`{"number1":100,"number2":200}`)
	data1 := Data{}
	json.Unmarshal(b1, &data1)
	fmt.Printf("data1.Number1: %d\n", data1.Number1)
	fmt.Printf("data1.Number2: %d\n", *data1.Number2)

	b2 := []byte("{}") //‥‥‥①
	data2 := Data{}
	json.Unmarshal(b2, &data2)
	fmt.Printf("data2.Number1: %d\n", data2.Number1) //‥‥‥②
	fmt.Printf("data2.Number2: %v\n", data2.Number2) //‥‥‥③
}
