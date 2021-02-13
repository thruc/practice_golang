package main

import (
	"fmt"
)

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var maxNum int
	for _, s := range l {
		if maxNum < s {
			maxNum = s
		}
	}
	fmt.Println(maxNum)
	var sumNum int

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	for _, v := range m {
		sumNum += v
	}
	fmt.Println(sumNum)

}
