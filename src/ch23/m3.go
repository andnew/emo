package ch23

import "fmt"

// 第三个案例 死锁
func Demo3() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 总结
	// 创建了一个缓冲信道，其容量为 2。由于该信道的容量为 2，因此可向它写入两个字符串，而且不会发生阻塞。
	// 在 ch <- "naveen" 和 ch <- "paul" ，我们向信道写入两个字符串，该信道并没有发生阻塞。
	// 但是 ch <- "steve" 写入发生阻塞 现在想要这次写操作能够进行下去，必须要有其它协程来读取这个信道的数据。
	// 在 ch <- "steve" 代码行 发生 死锁 现象
	// fatal error: all goroutines are asleep - deadlock!
}
