package ch29

import (
	"fmt"
	"sync"
)

// defer 的实际应用
// 目前为止，我们看到的代码示例，都没有体现出 defer 的实际用途。本节我们会看看 defer 的实际应用。

// 当一个函数应该在与当前代码流（Code Flow）无关的环境下调用时，可以使用 defer。
// 我们通过一个用到了 WaitGroup 代码示例来理解这句话的含义。我们首先会写一个没有使用 defer 的程序，然后我们会用 defer 来修改，看到 defer 带来的好处。

func (r rect) area1(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)

}

func MainCall6() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area1(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// 总结&分析
// 使用 defer 还有一个好处。假设我们使用 if 条件语句，又给 area 方法添加了一条返回路径（Return Path）。
// 如果没有使用 defer 来调用 wg.Done()，我们就得很小心了，确保在这条新添的返回路径里调用了 wg.Done()。
// 由于现在我们延迟调用了 wg.Done()，因此无需再为这条新的返回路径添加 wg.Done() 了。
