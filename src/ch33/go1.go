package ch33

import "fmt"

func MainCall1() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)
}

// 总结&分析
// 支持头等函数（First Class Function）的编程语言，可以把函数赋值给变量，也可以把函数作为其它函数的参数或者返回值。Go 语言支持头等函数的机制。
// 变量a 是一个函数，a() 表示 调用函数
