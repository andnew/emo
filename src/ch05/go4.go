package ch05

import "fmt"

// 布尔常量
// 布尔常量和字符串常量没有什么不同。他们是两个无类型的常量 true 和 false。
// 字符串常量的规则适用于布尔常量，所以在这里我们不再重复。

func MainCall7()  {
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst // 不允许
	var customBool myBool = trueConst // 允许
	//defaultBool = customBool // 不允许
	fmt.Println(defaultBool, customBool)
}
// 总结&分析
//