package ch29

import (
	"fmt"
	"sync"
)

// defer 的实际应用
// 目前为止，我们看到的代码示例，都没有体现出 defer 的实际用途。本节我们会看看 defer 的实际应用。

// 当一个函数应该在与当前代码流（Code Flow）无关的环境下调用时，可以使用 defer。
// 我们通过一个用到了 WaitGroup 代码示例来理解这句话的含义。我们首先会写一个没有使用 defer 的程序，然后我们会用 defer 来修改，看到 defer 带来的好处。

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}

func MainCall5() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// 总结&分析
// wg.Done() 代码 重复出现了，采用 defer  进行处理 见 go6 示例
