package ch33

import "fmt"

//用户自定义的函数类型

type add func(a int, b int) int

func MainCall4() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}

// func 也是一种类型
