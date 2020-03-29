package ch11

import "fmt"

//使用 range 迭代数组

func MainCall10() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}

// for 循环遍历数组的元素, 从 索引 0 到 数组长度 length - 1
