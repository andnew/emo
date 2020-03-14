package ch23

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup 是一个结构体类型，
// 我们在第 18 行创建了 WaitGroup 类型的变量，其初始值为零值。
// WaitGroup 使用计数器来工作。
// 当我们调用 WaitGroup 的 Add 并传递一个 int 时，WaitGroup 的计数器会加上 Add 的传参。
// 要减少计数器，可以调用 WaitGroup 的 Done() 方法。
// Wait() 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。

func Demo5(i int, wg *sync.WaitGroup) {

	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
