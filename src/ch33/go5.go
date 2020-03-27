package ch33

import "fmt"

// 把高阶函数（Hiher-order Function）定义为：满足下列条件之一的函数：

// 接收一个或多个函数作为参数
// 返回值是一个函数

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func MainCall5() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}

// 总结&分析
// 定义了一个函数 simple，simple 接收一个函数参数（该函数接收两个 int 参数，返回一个 a 整型）。
// 在 MainCall5 函数中, 创建了一个匿名函数 f，其签名符合 simple 函数的参数。我们在下一行调用了 simple，并传递了参数 f。
// 该程序打印输出 67。
