package ch04

import "fmt"

//下面是 Go 支持的基本类型：
// bool
// 数字类型
//		int8, int16, int32, int64, int
//		uint8, uint16, uint32, uint64, uint
//		float32, float64
//		complex64, complex128
//		byte
//		rune
//	string

// bool 类型表示一个布尔值，值为 true 或者 false
func MainCall1() {
	a := true
	b := false
	fmt.Println("Func Main Call 1 , a:", a, "b:", b)
	c := a && b
	fmt.Println("Func Main Call 1 , c:", c)
	d := a || b
	fmt.Println("Func Main Call 1 , d:", d)
}

// 总结&分析
// bool 的值 有2个 true 和 false
// 简短赋值、运行时赋值 应用
