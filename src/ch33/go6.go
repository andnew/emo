package ch33

import "fmt"

func simple6() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func MainCall6() {
	s := simple6()
	fmt.Println(s(60, 7))
}

// 总结&分析
// simple 返回的是一个函数类型
