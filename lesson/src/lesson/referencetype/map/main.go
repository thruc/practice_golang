package main

import "fmt"

// 辞書のような働き
func main() {
	// 明示
	var m = map[string]int{"a": 100, "b": 200}
	fmt.Println(m)

	// 暗示
	m2 := map[string]int{"a": 100, "b": 200}
	fmt.Println(m2)

	// 改行しても大丈夫
	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)
	// meke関数による生成
	m4 := make(map[int]string)
	m4[1] = "US"
	m4[2] = "JAPAN"
	m4[3] = "CHINA"
	fmt.Println(m4)
	fmt.Println(m4[5]) //　定義さ入れていないと空文字（初期値）

	s := m["a"]
	fmt.Println(s)

	s1 := m["c"]
	fmt.Println(s1)

	s3, ok := m["b"]
	fmt.Println(s3)
	// エラーハンドリング
	if !ok {
		fmt.Println("error")
	}

	s4, ok := m["c"]
	//_, ok = m["b"]
	fmt.Println(ok)

	fmt.Println(s4)

	m["b"] = 300

	m["c"] = 500

	// 削除
	delete(m, "c")

	fmt.Println(len(m))

	// for
	for k, v := range m {
		fmt.Println(k, v)
	}

	m5 := map[int][]string{
		1: {"A"},
		2: {"B", "C"},
		3: {"D", "E"},
	}
	fmt.Println(m5)

	m6 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}

	fmt.Println(m6)

	// panic: assignment to entry in nil map
	// var m7 map[string]int
	// m7["pc"] = 5000
	// fmt.Println(m7)

}
