package ch11

import "fmt"

// Go 中的数组是值类型而不是引用类型
func MainCall7() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

// 总结&分析
