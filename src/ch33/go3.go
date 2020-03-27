package ch33

import "fmt"

func MainCall3() {
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}

// 函数() 直接调用
