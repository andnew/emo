package ch11

import "fmt"

// 切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用

func MainCall13() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(b)
}

// 总结&分析
// 语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片
