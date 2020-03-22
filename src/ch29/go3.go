package ch29

import "fmt"

// 实参取值（Arguments Evaluation）

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
func MainCall3() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

}

// 总结&分析
// 在 Go 语言中，并非在调用延迟函数的时候才确定实参，而是当执行 defer 语句的时候，就会对延迟函数的实参进行求值
// 代码运行效果
// value of a before deferred function call 10
// value of a in deferred function 5
