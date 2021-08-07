package main

import "fmt"

func main() {
	var a []int = []int{100, 200}
	a = append(a, 300) // 要素数が変更できるので要素を追加できる
	fmt.Println(a)

	m := make([]int, 5) // int型初期値５つを作成
	fmt.Println(m)

	l := make([]int, 1, 5) // int型５つを作成
	fmt.Println(l)

	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)

	l = append(l, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)

	l = append(l, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)

	e := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(e), cap(e), e)

}
