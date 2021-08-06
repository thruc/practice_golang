package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	ffl64 := 3.2
	fmt.Printf("%T\n", ffl64)

	// 基本定期にはf３２は使用しない
	var fl32 float32 = 5.4
	fmt.Printf("%T\n", fl32)

	// f32に変換
	fmt.Printf("%T\n", float32(ffl64))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	// 正の無限大
	ninf := -1.0 / zero
	fmt.Println(ninf)

	// 負の目元代
	nan := zero / zero
	fmt.Println(nan)

	//var u8 uint = 255 byte

	var c64 complex64 = -5 + 12i

	//【uint型(+の整数）、complex(複素数型)】
	//他にもuint(+の整数) complex(複素数型）などの型があるが、あまり使用頻度は無いので今回は紹介だけにします。

	var u8 uint = 255 //uint型

	fmt.Println(u8)
	fmt.Printf("%T\n", u8)

	fmt.Println(c64)
	fmt.Printf("%T\n", c64)

}
