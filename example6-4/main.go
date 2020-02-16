package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type NullInt struct { //‥‥‥①
	Int   int
	Valid bool
}

func (n *NullInt) UnmarshalJSON(b []byte) error { //‥‥‥②
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	n.Int = int(i)
	n.Valid = true
	return nil
}

type Data struct {
	Number1 int     `json:"number1"`
	Number2 NullInt `json:"number2"` //‥‥‥③
}

func main() {
	b1 := []byte(`{"number1":100,"number2":200}`)
	data1 := Data{}
	json.Unmarshal(b1, &data1)
	fmt.Printf("data1.Number1: %d\n", data1.Number1)
	fmt.Printf("data1.Number2: %v\n", data1.Number2) //‥‥‥④

	b2 := []byte("{}")
	data2 := Data{}
	json.Unmarshal(b2, &data2)
	fmt.Printf("data2.Number1: %d\n", data2.Number1)
	fmt.Printf("data2.Number2: %v\n", data2.Number2) //‥‥‥⑤

	if data1.Number2.Valid { //‥‥‥⑥
		fmt.Printf("data1.Number2 value: %d\n",
			data1.Number2.Int)
	} else {
		fmt.Println("data1.Number2 is null")
	}

	if data2.Number2.Valid { //‥‥‥⑦
		fmt.Printf("data2.Number2 value: %d\n",
			data2.Number2.Int)
	} else {
		fmt.Println("data2.Number2 is null")
	}
}
