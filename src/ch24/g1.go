package ch24

import (
	"fmt"
	"time"
)

// select 语句用于在多个发送/接收信道操作中进行选择。
// select 语句会一直阻塞，直到发送/接收操作准备就绪。如果有多个信道操作准备完毕，select 会随机的选取其中之一执行。
// 该语法与 switch 类似，所不同的是，这里的每个 case 语句都是信道操作。

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "From Server1 ...."
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "From Server2 ...."
}

// main 函数代码
func MainCall1() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

// 总结 select 更加简洁的，替换之前的执行完成的信道 close 操作
// 函数 server1 休眠了 6 秒，接着将 文本 From Server1 写入信道 ch.
// 函数 server2 休眠了 3 秒，然后将 文本 From Server2 写入信道 ch
// 函数 MainCall1 代码中 运行到 select 语句。 select 会一直发生阻塞，除非其中有 case 准备就绪。
// 在上述程序里， server1 协程会在 6 秒之后写入 output1 信道,而 server2 协程会在 3 秒之后写入 output2 信道。
// 因此 select 语句会阻塞 3 秒钟，等着 server2 向 output2 信道写入数据。 3秒钟过后，程序就会输出
// From Server2 ....
