package ch04

import "fmt"

// complex64：实部和虚部都是 float32 类型的的复数。
// complex128：实部和虚部都是 float64 类型的的复数。

// 内建函数 complex 用于创建一个包含实部和虚部的复数。complex 函数的定义如下：
// func complex(r, i FloatType) ComplexType

// 该函数的参数分别是实部和虚部，并返回一个复数类型。实部和虚部应该是相同类型，也就是 float32 或 float64。
// 如果实部和虚部都是 float32 类型，则函数会返回一个 complex64 类型的复数。如果实部和虚部都是 float64 类型，则函数会返回一个 complex128 类型的复数。

func MainCall5() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("Func Main Call 5 , sum:", cadd)
	cmul := c1 * c2
	fmt.Println("Func Main Call 5 , product:", cmul)
}
