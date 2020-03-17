package ch05

import (
	"fmt"
	"math"
)

// 在 Go 语言中，术语"常量"用于表示固定的值。比如 5 、-89、 I love Go、67.89 等等。

func MainCall1() {
	var a int = 50
	var b string = "I love Go "

	fmt.Println("a = ", a, " , b = ", b)

	a = 100
	b = "love Go too"
	fmt.Println("a = ", a, " , b = ", b)
}

// 总结&分析
// 变量 a 和 b 分别被赋值为常量 50 和 I love GO
// 关键字 const 被用于表示常量，比如 50 和 I love Go。即使在上面的代码中我们没有明确的使用关键字 const，但是在 Go 的内部，它们是常量。
// a 和 b 成为 变量， 是存储特定类型的值的单元

// const 关键词修饰的变量 a  不可以重复被赋值
func MainCall2() {
	const a = 55 // 允许
	//a = 89       // 不允许重新赋值
	fmt.Printf("a = %d", a)
}

//总结&分析
//编译出错,报错信息是 Cannot assign to a

// 常量的值会在编译的时候确定。因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量。
func MainCall3() {
	fmt.Println("Hello, playground")
	var a = math.Sqrt(4) // 允许
	//const b = math.Sqrt(4) // 不允许
	fmt.Println("a = ", a)
}

// 总结&分析
// 变量可以在运行时进行赋值,常量必须在初始化是进行赋值
// 编译报错 Const initializer 'math.Sqrt(4)' is not a constant

