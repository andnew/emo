package ch23

import "fmt"

// 第一个案例
func Demo1() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 总结
	// 创建了一个缓冲信道，其容量为 2。由于该信道的容量为 2，因此可向它写入两个字符串，而且不会发生阻塞。
	// 在 ch <- "naveen" 和 ch <- "paul" ，我们向信道写入两个字符串，该信道并没有发生阻塞。
	// 我们又在 fmt.Println(<-ch) 和 fmt.Println(<-ch) 分别读取了这两个字符串。
}
