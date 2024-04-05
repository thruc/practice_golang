package main

import "fmt"

type customer struct {
	id         string
	operations []float64
}

func main() {
	var cust1 any = customer{id: "x", operations: []float64{1.0, 2.0}}
	var cust2 any = customer{id: "x", operations: []float64{1.0, 2.0}}
	fmt.Println(cust1 == cust2) // コンパイルはできるが実行時error
	/*
		panic: runtime error: comparing uncomparable type main.customer
	*/
}
