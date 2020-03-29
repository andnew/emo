package ch11

import "fmt"

// 切片是动态的，使用 append 可以将新元素追加到切片上。append 函数的定义是 func append（s[]T，x ... T）[]T。
//x ... T 在函数定义中表示该函数接受参数 x 的个数是可变的。这些类型的函数被称为可变函数。
func MainCall20() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	fmt.Println("前地址:", &cars)
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6
	fmt.Println("后地址:", &cars)

}
