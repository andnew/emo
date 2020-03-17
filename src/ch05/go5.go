package ch05

import "fmt"

//数字常量包含整数、浮点数和复数的常量。

func MainCall8() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}

// 总结&分析
// 常量 a 是没有类型的，它的值是 5
// 下面的变量 使用 a 进行赋值，相应 a 转化 对应的数据类型
