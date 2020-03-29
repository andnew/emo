package ch11

import "fmt"

// 通过将数组作为参数传递给 len 函数，可以得到数组的长度。

func MainCall9() {
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a))
}
