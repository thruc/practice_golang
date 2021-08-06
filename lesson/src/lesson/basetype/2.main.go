package main

import (
	"fmt"
	"math"
)

func main() {

	var i int = 100

	fmt.Println(i + 50)

	var ii int64 = 50

	//fmt.Println(i + ii)

	fmt.Println(int32(ii))
	fmt.Printf("%T\n", int32(ii))
	fmt.Printf("int8   -> min:%d, max:%d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16  -> min:%d, max:%d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32  -> min:%d, max:%d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64  -> min:%d, max:%d\n", int64(math.MinInt64), int64(math.MaxInt64))

	fmt.Printf("uint8  -> min:0, max:%d\n", math.MaxUint8)          // uint8の最小値は0。最小値の定数は無い
	fmt.Printf("uint16 -> min:0, max:%d\n", math.MaxUint16)         // uint16の最小値は0。最小値の定数は無い
	fmt.Printf("uint32 -> min:0, max:%d\n", uint32(math.MaxUint32)) // uint32の最小値は0。最小値の定数は無い
	fmt.Printf("uint64 -> min:0, max:%d\n", uint64(math.MaxUint64)) // uint64の最小値は0。最小値の定数は無い
	val1 := 123
	fmt.Printf("val1 -> %d, type: %T\n", val1, val1)

	val2 := int(123)
	fmt.Printf("val2 -> %d, type: %T\n", val2, val2)

	val3 := uint(123)
	fmt.Printf("val3 -> %d, type: %T\n", val3, val3)

	val4 := int16(123)
	fmt.Printf("val4 -> %d, type: %T\n", val4, val4)

	val5 := int32(123)
	fmt.Printf("val5 -> %d, type: %T\n", val5, val5)

	val6 := int64(123)
	fmt.Printf("val6 -> %d, type: %T\n", val6, val6)
}
