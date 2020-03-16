package ch03

import (
	"fmt"
	"math"
)

//第三章 变量
//变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值。

// 语法
// var name type 是声明单个变量的语法

func MainCall1() {
	var age int // 声明变量 age 是 int 类型
	fmt.Println("Func mainCall1 , my age is ", age)
	age = 29
	fmt.Println("Func mainCall1 , my age is ", age)
	age = 54
	fmt.Println("Func mainCall1 , my new age is ", age)
}

// 总结
// 变量 age 声明 var age int
// age = 29 赋值语句
// age = 54 二次赋值，覆盖掉了之前的 29

func MainCall2() {
	var age int = 29
	fmt.Println("Func mainCall2 , my age is ", age)
}

// 总结&分析
// var age int = 29 声明、赋值 一行代码 搞定

// 类型推断（Type Inference）
// 如果变量有初始值，那么 Go 能够自动推断具有初始值的变量的类型。因此，如果变量有初始值，就可以在变量声明中省略 type。
// 如果变量声明的语法是 var name = initialvalue，Go 能够根据初始值自动推断变量的类型。
func MainCall3() {
	var age = 29
	fmt.Println("Func mainCall3 , my age is ", age)
}

// 总结&分析
// var age = 29 进行类型推算 从而 推断出 age 是 int 类型

// 声明多个变量
// Go 能够通过一条语句声明多个变量。
//声明多个变量的语法是 var name1, name2 type = initialvalue1, initialvalue2。
func MainCall4() {
	var width, height int = 100, 50 // 声明多个变量
	fmt.Println("Func mainCall4 , width is", width, "height is", height)
}

// 总结&分析
// 一行声明多个变量并且赋值

func MainCall5() {
	var width, height int
	fmt.Println("Func mainCall5 , width is", width, "height is", height)
	width = 100
	height = 50
	fmt.Println("Func mainCall5 , new width is", width, "new height is ", height)
}

// 总结&分析
// var width, height int 没有初始化值，那么 int 的 ，默认是 0

// 在有些情况下，我们可能会想要在一个语句中声明不同类型的变量。其语法如下
// var (
//   name1 = initialvalue1
//   name2 = initialvalue2
//   ....
// )
func MainCall6() {
	var (
		name   = "naveen"
		age    = 29
		height int
	)
	fmt.Println("Func mainCall6 , my name is", name, ", age is", age, "and height is", height)
}

// 总结&分析
// 不同类型的变量 var 括号内 分行 声明/初始化

// 简短声明
// Go 也支持一种声明变量的简洁形式，称为简短声明（Short Hand Declaration），该声明使用了 := 操作符。
// 声明变量的简短语法是 name := initialvalue。
func MainCall7() {
	name, age := "naveen", 29 // 简短声明

	fmt.Println("Func mainCall7 , my name is", name, "age is", age)
}

// 总结&分析
// 简短的声明，没有关键字 var ,直接采用 类型推导 的方式，变量数 必须与 参数值个数一致

//func MainCall8() {
//	name, age := "naveen" //error
//
//	fmt.Println("my name is", name, "age is", age)
//}
// 总结&分析
// Assignment count mismatch: 2 = 1

// 简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的。
func MainCall9() {
	a, b := 20, 30 // 声明变量a和b
	fmt.Println("Func mainCall9 , a is", a, "b is", b)
	b, c := 40, 50 // b已经声明，但c尚未声明
	fmt.Println("Func mainCall9 , b is", b, "c is", c)
	b, c = 80, 90 // 给已经声明的变量b和c赋新值
	fmt.Println("Func mainCall9 , changed b is", b, "c is", c)
}

// 总结&分析
// 代码执行结果如下
// Func mainCall9 , a is 20 b is 30
// Func mainCall9 , b is 40 c is 50
// Func mainCall9 , changed b is 80 c is 90

//
func MainCall10() {
	a, b := 20, 30 // 声明a和b
	fmt.Println("a is", a, "b is", b)
	//a, b := 40, 50 // 错误，没有尚未声明的变量
}

// 总结&分析
// 上面运行后会抛出 No new variables on left side of := 的错误，这是因为 a 和 b 的变量已经声明过了，:= 的左边并没有尚未声明的变量。

// 变量也可以在运行时进行赋值
func MainCall11() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Func mainCall11, minimum value is ", c)
}

// 总结&分析
// 简短声明，是可以在运行进行赋值的

// 由于 Go 是强类型（Strongly Typed）语言，因此不允许某一类型的变量赋值为其他类型的值。
// 下面的程序会抛出错误 cannot use "naveen" (type string) as type int in assignment，
// 这是因为 age 本来声明为 int 类型，而我们却尝试给它赋字符串类型的值。
func MainCall12() {
	age := 29      // age是int类型
	age = "naveen" // 错误，尝试赋值一个字符串给int类型变量
}

// 总结&分析
// Cannot use '"naveen"' (type string) as type int in assignment
