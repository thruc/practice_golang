package main

func main() {
	//Goの無名関数はクロージャーで、クロージャーとは日本語では関数閉包と呼ばれ、関数と関数の処理に関する関数外の環境をセットして閉じ込めた物です。
	f := func(x, y int) int {
		return x + y
	}
	f(1, 2)

	func(x, y int) int {
		return x + y
	}(1, 2)

}
