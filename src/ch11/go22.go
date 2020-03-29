package ch11

import "fmt"

// 可以使用 ... 运算符将一个切片添加到另一个切片。

func MainCall22() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}
