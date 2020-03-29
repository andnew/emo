package ch11

import "fmt"

func MainCall5() {
	a := [...]int{12, 78, 50}
	fmt.Println(a)
}

// 总结&分析
// 声明时忽略数组的长度，并用 ... 代替，让编译器自动计算长度
