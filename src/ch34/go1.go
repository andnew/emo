package ch34

import "fmt"

// 反射就是程序能够在运行时检查变量和值，求出它们的类型

func MainCall1() {

	i := 10
	fmt.Printf("%d %T \n", i, i)
}

// 总结&分析
// 通过 Printf 的 方式 获取 数据类型
