package ch11

import "fmt"

func MainCall2() {
	var a [3]int //int array with length 3
	a[0] = 12    // array index start at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}

// 总结&分析
// 数组的索引 从 0 开始 到 length - 1 结束。
