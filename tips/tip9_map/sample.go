package main

import "fmt"

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	fmt.Println(m)
	// 全く結果の予想ができない

	// コピーをすることで冪等性を保証する
	m2 := copyMap(m)
	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}
	fmt.Println(m2)
}

func copyMap[K comparable, T any](m map[K]T) map[K]T {
	m2 := make(map[K]T, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return m2
}
