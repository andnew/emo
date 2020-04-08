package ch32

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
		debug.PrintStack()
	}
}

func aa() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func MainCall4() {
	aa()
	fmt.Println("normally returned from MainCall7")
}

// 总结&分析
// 数组越界这种错误，引发panic 同样可以采取recover的收集错误
