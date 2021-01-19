package main

import (
	"fmt"
)

//関数の外で宣言できる
var i int = 1
var f64 float64 = 1.2
var s string = "str"
var t, f bool = true, false

const (
	username string = "test_user"
	password string = "test_password"
)

func constFoo() {
	fmt.Println(username, password)

}
func foo() {
	//関数内のみショートを使って宣言できる
	xi := 1
	xf64 := 1.2
	xs := "str"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	//型を調べる
	fmt.Printf("%T", f64)

	fmt.Println()
	foo()
	constFoo()

}
