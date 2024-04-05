package main

import "fmt"

func main() {
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))
	/*
		position 0: h
		position 1: Ã
		position 3: l
		position 4: l
		position 5: o
		len=6
	*/

	fmt.Println("----")
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}

	fmt.Println("----")
	for i, r := range []rune(s) {
		fmt.Printf("position %d: %c\n", i, r)
	}
}
