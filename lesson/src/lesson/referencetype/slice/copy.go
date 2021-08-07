package main

import "fmt"

func main() {
	var a []int = []int{100, 200}
	b := a // 参照渡し
	fmt.Println(a)

	b[0] = 1000 // 参照渡しなのでaの値も変更される

	fmt.Println(a)

	d := []int{1, 2, 3, 4, 5, 6}

	copy := copy(d, a)
	a[0] = 1000
	fmt.Println(copy) // コピーに成功した要素数
	fmt.Println(d)
}
