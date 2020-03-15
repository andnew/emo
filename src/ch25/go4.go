package ch25

import (
	"fmt"
	"sync"
)

// 使用信道来处理竟态条件

var z = 0

func increment3(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}

func MainCall4() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment3(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of z ", z)
}
// 总结&分析
// 上述程序中，ch := make(chan bool, 1) 创建了容量为1 的 缓冲信道，
// 通过代码行 go increment3(&w, ch) 将 ch 传入 increment3 协程。
// 该缓冲信道用于保证只有一个协程访问增加 x 的临界区。
// 具体的实现方法是在 z 增加之前（ z = z + 1 行），传入 true 给缓冲信道。由于缓冲信道的容量为 1，
// 所以任何其他协程试图写入该信道时，都会发生阻塞，直到 z 增加后，信道的值才会被读取（第 <-ch 行）。
// 实际上这就保证了只允许一个协程访问临界区。

