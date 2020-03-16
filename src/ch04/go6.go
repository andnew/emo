package ch04

import "fmt"

// string 类型
// 在 Golang 中，字符串是字节的集合。

func MainCall6() {
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("Func Main Call 6 , My name is", name)
}

// 总结&分析
// 字符串的拼接 最简单的方式 采用 + 操作
