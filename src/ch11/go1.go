package ch11

import "fmt"

// 数组的声明

func MainCall1() {

	var a [3]int // int array with length 3
	fmt.Println(a)
}

// 总结&分析
// 变量 a 是 一个长度为3的整型数组；数组中的所有元素都自动赋值为 数组类型的零值
