package ch25

import (
	"fmt"
	"sync"
)

//Mutex
//Mutex 用于提供一种加锁机制（Locking Mechanism），可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。
//
//Mutex 可以在 sync 包内找到。Mutex 定义了两个方法：Lock 和 Unlock。所有在 Lock 和 Unlock 之间的代码，都只能由一个 Go 协程执行，于是就可以避免竞态条件。
//
//mutex.Lock()
//x = x + 1
//mutex.Unlock()
//在上面的代码中，x = x + 1 只能由一个 Go 协程执行，因此避免了竞态条件。
//
//如果有一个 Go 协程已经持有了锁（Lock），当其他协程试图获得该锁时，这些协程会被阻塞，直到 Mutex 解除锁定为止。

// 第二个示例

// 共享变狼 x
var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func MainCall2() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("Final value of x", x)
}
// 总结&分析
// 函数 increment 代码中 把 x 的值加 1 ，并调用 WaitGroup 的 Done(),通知该函数已经结束。
// 程序中，生成了 1000 个 increment 协程。每个Go协程并发的运行， 由于 x = x + 1 试图增加 x 的值，
// 因此 多个并发的协程试图访问 x 的值。 这时就会发生竟态条件。

