package ch24

import "fmt"

// 如果 select 只含有值为 nil 的信道，也同样会执行默认情况

func MainCall5() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value ", v)
	default:
		fmt.Println("default case executed")
	}
}
// 总结&分析
// 在上面的程序中，ch 等于 nil
