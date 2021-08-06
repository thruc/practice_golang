package main

import "fmt"

func main() {

	// 配列　他言語と違いは初期値の要素数を変更できない
	var arr1 [2]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	// 三つ目には初期値のから文字が入る
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数を自動で数えてくれる
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	var arr5 [5]int = [5]int{1, 2, 3}
	fmt.Println(arr5)

	// 値の取り出し
	fmt.Println(arr4[0])
	fmt.Println(arr4[2-1])

	// 代入
	arr4[1] = "F"
	fmt.Println(arr4)

	// エラー
	//fmt.Println(arr1 == arr5)
	//fmt.Println(arr2 == arr3)

	fmt.Println(len(arr1))

}
