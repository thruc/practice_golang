package main

import "fmt"

// 明示的に宣言
// var 変数名 型 = 値
var i int = 100
var s string = "Golang"

// 複数宣言
var t, f bool = true, false

var (
	ii int    = 1000
	ss string = "Go"
)

// 型のみ宣言
// intの初期値は0
var i2 int

// stringの初期値はから文字
var sss string

// 暗黙的な宣言は関数外だとエラー
//i5 := 100

// 明示的に宣言と暗黙的な宣言の使い分けは関数の外か中か
// 可能であれば明示的な宣言を行う方が良い

func main() {
	fmt.Println(i2)
	i2 = 150
	fmt.Println(i2)

	fmt.Println(sss)
	sss = "Go!!"
	fmt.Println(sss)

	i2 = 200
	fmt.Println(i2)

	// 暗黙的な宣言
	// 変数名 := 値
	i3 := 200
	fmt.Println(i3)

	// 再宣言はエラー
	//i3 := 300

	i3 = 300

	// 型が違うとエラー
	//i3 = "string"

	// 暗黙的な宣言は関数外だとエラー
	//fmt.Println(i5)

}
