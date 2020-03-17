package ch05

import "fmt"

// Go 的类型策略不允许将一种类型的变量赋值给另一种类型的变量

func MainCall6() {
	var defaultName = "Sam" // 允许
	type myString string
	var customName myString = "Sam" // 允许
	//customName = defaultName        // 不允许
	fmt.Printf("defaultName 的类型是 %T , 值是 %v\n", defaultName, defaultName)
	fmt.Printf("customName 的类型是 %T , 值是 %v\n", customName, customName)
}
// 总结&分析
// 1、代码行 customName = defaultName   存在的情况下
//    执行结果是
//    cannot use defaultName (type string) as type myString in assignment

// 2、代码行 customName = defaultName 注释掉的时候，运行结果是
//    defaultName 的类型是 string , 值是 Sam
//    customName 的类型是 ch05.myString , 值是 Sam