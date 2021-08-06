package main

import (
	"fmt"
)

// 関数を引数にとる関数
func callFunction(f func()) {
	f()
}

func main() {
	f := func() {
		fmt.Println("I'm a function")
	}
	callFunction(f) // => "I'm a function"
}
