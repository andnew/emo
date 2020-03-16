package ch04

import (
	"fmt"
	"unsafe"
)

// 有符号整型
//	int8：表示 8 位有符号整型
//	大小：8 位
//	范围：-128～127

//  int16：表示 16 位有符号整型
//	大小：16 位
//  范围：-32768～32767

//   int32：表示 32 位有符号整型
//   大小：32 位
//   范围：-2147483648～2147483647

//   int64：表示 64 位有符号整型
//   大小：64 位
//   范围：-9223372036854775808～9223372036854775807

//   int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
//	 大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
//	 范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807。

func MainCall2() {
	var a int = 89
	b := 95
	fmt.Println("Func Main Call 2 , value of a is", a, "and b is", b)
}

// 总结&分析
// var a int = 89 标准的声明 + 赋值 方式
// b := 95 采用的是 简短赋值 方式， 推测 b 的 为 int , int 根据 机器 32 / 64 进行区分 为 32 还是 64 位的

// 在 Printf 方法中，使用 %T 格式说明符（Format Specifier），可以打印出变量的类型。
// Go 的 unsafe 包提供了一个 Sizeof 函数，该函数接收变量并返回它的字节大小。
// unsafe 包应该小心使用，因为使用 unsafe 包可能会带来可移植性问题。
func MainCall3() {
	var a int = 89
	b := 95
	fmt.Println("Func Main Call 3 , value of a is", a, "and b is", b)
	fmt.Printf("Func Main Call 3 , type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   // a 的类型和大小
	fmt.Printf("\nFunc Main Call 3 , type of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
}

// 总结&分析
// 输出结果是
// Func Main Call 3 , value of a is 89 and b is 95
// Func Main Call 3 , type of a is int, size of a is 8
// Func Main Call 3 , type of b is int, size of b is 8
// 从输出中可以推断出 a 和 b 为 int 类型，且大小都是 64 位（8 字节）。如果你在 32 位系统上运行上面的代码，会有不同的输出。在 32 位系统下，a 和 b 会占用 32 位（4 字节）的大小
