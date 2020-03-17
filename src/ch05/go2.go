package ch05

import "fmt"

// 字符串常量
// 双引号中的任何值都是 Go 中的字符串常量。例如像 Hello World 或 Sam 等字符串在 Go 中都是常量。
// 什么类型的字符串属于常量？答案是他们是无类型的。

// 像 Hello World 这样的字符串常量没有任何类型。
// const hello = "Hello World"
func MainCall4()  {
	const hello = "Hello World"
	fmt.Printf("type %T value %v\n", hello, hello)
}
// 总结&分析
// 执行结果是
// type string value Hello World

func MainCall5()  {
	var name = "Sam"
	fmt.Printf("type %T value %v\n", name, name)
}
// 总结&分析
// type string value Sam


