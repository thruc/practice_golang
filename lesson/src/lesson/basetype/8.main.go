package main

import (
	"fmt"
	"strconv"
)

func main() {

	//　型変換
	fl64 := 10.2
	fmt.Printf("%T \n", fl64)

	n := int(fl64)
	fmt.Printf("%T \n", n)
	nn := float64(n)
	fmt.Printf("%T \n", nn)

	var s string = "100"

	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))

	var ii int = 321
	var ss string
	ss = strconv.Itoa(ii)
	fmt.Println(ss)

}
