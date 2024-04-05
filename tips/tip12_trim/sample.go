package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("123oxo", "xo"))
	// 123
	fmt.Println(strings.TrimSuffix("123oxo", "xo"))
	// 123o
}
