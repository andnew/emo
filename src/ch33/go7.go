package ch33

import "fmt"

// 闭包（Closure）是匿名函数的一个特例。当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。

func MainCall7() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

// 总结&分析
// func()函数 访问变量 a
