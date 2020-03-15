package ch24

import "fmt"

// 使用默认情况

func MainCall4() {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
}
// 总结&分析
// 程序执行结果 default case executed