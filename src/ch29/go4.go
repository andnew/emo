package ch29

import "fmt"

// defer 栈
// 当一个函数内多次调用 defer 时，Go 会把 defer 调用放入到一个栈中，随后按照后进先出（Last In First Out, LIFO）的顺序执行。

func MainCall4() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: \n")
	for i, v := range []rune(name) {
		defer fmt.Printf("------MainCall4----%d----%c\n", i, v)
	}
}

// 总结&分析
// 代码运行效果
// Orignal String: Naveen
// Reversed String:
// ------MainCall4----5----n
// ------MainCall4----4----e
// ------MainCall4----3----e
// ------MainCall4----2----v
// ------MainCall4----1----a
// ------MainCall4----0----N
