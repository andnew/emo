package ch23

import "fmt"

// 第二个案例
func Demo2(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, " to ch ")
	}
	close(ch)
}

// 总结
// 声明一个信道 ch ，容量 2 ，
// 通过 for 循环 依次的 ch 中写入数据，但是 ch 的容量是 2 ，在写入 第二个元素 后，ch 缓存已满，信道处于 阻塞 状态
