package ch04

import "fmt"

// 无符号整型
//	uint8：表示 8 位无符号整型
//	大小：8 位
//	范围：0～255

//	uint16：表示 16 位无符号整型
//	大小：16 位
//	范围：0～65535

//	uint32：表示 32 位无符号整型
//	大小：32 位
//	范围：0～4294967295

//	uint64：表示 64 位无符号整型
//	大小：64 位
//	范围：0～18446744073709551615

//	uint：根据不同的底层平台，表示 32 或 64 位无符号整型。
//	大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
//	范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。

//	浮点型
//	float32：32 位浮点数
//	float64：64 位浮点数

func MainCall4() {
	a, b := 5.67, 8.97
	fmt.Printf("\nFunc Main Call 4 , type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("Func Main Call 4 , sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("Func Main Call 4 , sum", no1+no2, "diff", no1-no2)
}

// 总结&分析
// a b 数据类型是 float64（float64 是浮点数的默认类型），注意 减法 精度位数
// Func Main Call 4 , type of a float64 b float64
// Func Main Call 4 , sum 14.64 diff -3.3000000000000007
// Func Main Call 4 , sum 145 diff -33
