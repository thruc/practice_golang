package main

import "fmt"

type account struct {
	balance float32
}

func main() {
	accounts := []account{
		{balance: 100.0},
		{balance: 200.0},
		{balance: 300.0},
	}
	for _, a := range accounts {
		a.balance += 1000
	}
	fmt.Printf("%+v\n", accounts) // [{balance:100} {balance:200} {balance:300}]

	// indexを用いることでスライス更新ができる
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Printf("%+v\n", accounts) // [{balance:1100} {balance:1200} {balance:1300}]

	s := []int{0, 1, 2}
	for range s {
		s = append(s, 10)
	}
	fmt.Printf("%+v\n", s) // [0 1 2 10 10 10]

	// ss := []int{0, 1, 2}
	// for i := 0; i < len(ss); i++ {
	// 	ss = append(ss, 10)
	// 	fmt.Printf("%+v\n", ss)
	// }
	// fmt.Println("無限loopでここまでこない")

	ch1 := make(chan int, 3) // ←要素 0、1、2 を含めるための最初のチャネルを作成する
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()
	ch2 := make(chan int, 3) // ←要素 10、11、12を含めるための 2 つ目のチャネルを作成する
	go func() {
		ch2 <- 10
		ch2 <- 11
		ch2 <- 12
		close(ch2)
	}()
	ch := ch1 // ←最初のチャネルをch へ代入する
	for v := range ch {
		// ← chを反復処理することで、消費チャネルを作成する
		fmt.Println(v)
		ch = ch2 // ← 2 つ目のチャネルを chに代入する
	}
	// 0
	// 1
	// 2

	a := [3]int{0, 1, 2}
	for i, v := range a {
		a[2] = 10 // vの値は更新されない
		fmt.Printf("%d %d\n", i, v)
		fmt.Printf("%+v\n", a)
		if i == 2 {
			fmt.Println(a[2])
		}
	}
	/*
		0 0
		[0 1 10]
		1 1
		[0 1 10]
		2 2
		[0 1 10]
		10
	*/
}
