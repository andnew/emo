package ch29

import "fmt"

func MainCall7() {

	fmt.Println("执行记录前==========")
	a := increaseA()
	fmt.Println("a=", a)
	b := increaseB()
	fmt.Println("b=", b)
	fmt.Println("执行记录后==========")

}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i //0
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r // 1
}

// 总结&分析
// 主要考察的是defer ， 返回值 带命名返回参数的函数
