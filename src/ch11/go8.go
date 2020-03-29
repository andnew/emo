package ch11

import "fmt"

// 当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func MainCall8() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}
