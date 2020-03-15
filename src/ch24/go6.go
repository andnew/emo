package ch24

import (
	"fmt"
	"time"
)

func server3(ch chan string) {
	ch <- "from server3"
}

func server4(ch chan string) {
	ch <- "from server4"
}

func MainCall6() {

	output1 := make(chan string)
	output2 := make(chan string)
	go server3(output1)
	go server4(output2)
	time.Sleep(1 * time.Second)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
// 总结&分析
// 程序在
//  go server3(output1)
//	go server4(output2)
// 代码行中分别被调用 server3 和 sever4 两个Go协程。接下来，主程序休眠 1 秒钟(time.Sleep(1 * time.Second) 代码行 )。
// 当程序控制到达 select 语句 server3 已经把 from server3 写到了 output1 信道上，而 server4 也同样把 from server4 写到了 output2 信道上。
// 因此这个 select 语句中的两种情况都准备好执行了。如果你运行这个程序很多次的话，输出会是 from server3 或者 from server4，这会根据随机选取的结果而变化。