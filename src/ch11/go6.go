package ch11

import "fmt"

func MainCall6() {
	a := [3]int{5, 78, 8}
	var b [3]int //保证正确,将数组长度修改为3
	b = a        // cannot use a (type [3]int) as type [5]int in assignment
	fmt.Println(b)
}

// 数组长度也是数组的定义的一部分
