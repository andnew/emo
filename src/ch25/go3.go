package ch25

import (
	"fmt"
	"sync"
)

// 使用Mutex
//

var y  = 0
func increment2(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}
func MainCall3() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of y", y)
}
// 总结&分析
// Mutex 是一个结构体类型。代码 var m sync.Mutex 创建了 Mutex 类型 m , 其值为领值。
// 修改 increment2 函数，将增加 x 的代码 (y = y + 1) 放置在 m.Lock() 和 m.Unlock() 之间。
// 这样这段代码不存在竟态条件了，因为任何时刻都只允许一个协程执行这段代码。

// 程序的执行结果 final value of y 1000
// go increment2(&w, &m) 代码行，传递 Mutex 的地址很重要.如果传递的是 Mutex 的值，而非地址，
// 那么每个协程都会得到 Mutex 的一份拷贝，竟态条件还是在发生。