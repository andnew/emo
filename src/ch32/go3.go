package ch32

import (
	"fmt"
	"time"
)

// 只有在相同的Go协程中调用recover才管用.
// recover 不能恢复一个不同协程的 panic

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	// 修改前 go b()
	b()
	time.Sleep(1 * time.Second)
}

func MainCall3() {
	a()
	fmt.Println("normally returned from main")
}

// 总结&分析
// 修改前 go b(), 导致 panic的发生与recover 不做相同的协程里。导致程序无法恢复
