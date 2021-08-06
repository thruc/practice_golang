package main

import "fmt"

func main() {

	// byte型

	// 配列で表現できる
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// 文字列に変更　アスキーコード
	fmt.Println(string(byteA))

	// byteに変換
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

}
