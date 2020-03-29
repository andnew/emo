package ch11

import "fmt"

// func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度。make 函数创建一个数组，并返回引用该数组的切片。

func MainCall19() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}
