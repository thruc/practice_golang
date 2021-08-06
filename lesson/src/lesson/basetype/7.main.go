package main

import "fmt"

func main() {

	// インターフェース
	var x interface{}

	// <nil> null的な存在
	fmt.Println(x)

	// 以下はすべて更新できるが型特有の計算ができない　すべての方と互換性を持つ
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [...]int{3, 4, 5, 6}
	fmt.Println(x)
	//x = 2
	//var y interface{} = 3
	//fmt.Println(x + y)
}
