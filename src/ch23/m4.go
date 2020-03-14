package ch23

import "fmt"

// 示例4 长度和容量
func Demo4() {

	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

// 总结
// ch := make(chan string, 3) 创建了一个容量为 3 的信道，于是它可以保存 3 个字符串。
// 接下来，
// 分别 ch <- "naveen" 行和 ch <- "paul" 行向信道写入了两个字符串。
// 于是 信道有两个字符串排队，因此其长度为 2 即 代码行 fmt.Println("length is", len(ch))的执行结果。
// 在 fmt.Println("read value", <-ch) 行，从信道读取了一个字符串。
// 现在该信道内只有一个字符串，因此其长度变为 1。
