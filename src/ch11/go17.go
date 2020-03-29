package ch11

import "fmt"

// 切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。

func MainCall17() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
}
