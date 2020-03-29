package ch11

import "fmt"

func MainCall14() {
	c := []int{6, 7, 8}
	fmt.Printf("切片=== %T , %v\n", c, c)

	b := [3]int{6, 7, 8}
	fmt.Printf("数组=== %T , %v\n", b, b)
}

// 总结&分析
// 切片的数据类型是 []int
// 数组的数据类型是 [length]int
