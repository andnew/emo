package ch11

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}
func MainCall23() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               // function modifies the slice
	fmt.Println("slice after function call", nos) // modifications are visible outside
}

// 总结&分析
// 切片在内部可由一个结构体类型表示。这是它的表现形式，
// type slice struct {
//     Length        int
//     Capacity      int
//     ZerothElement *byte
//  }

// 切片包含长度、容量和指向数组第零个元素的指针。
// 当切片传递给函数时，即使它通过值传递，指针变量也将引用相同的底层数组。
// 因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见。
