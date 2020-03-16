package ch04

import "fmt"

// 类型转换

// Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换。

func MainCall8() {
	//i := 55      //int
	//j := 67.8    //float64
	//sum := i + j //不允许 int + float64
	//fmt.Println(sum)
}

// 总结&分析
// 编译报错 Invalid operation: i + j (mismatched types int and float64)

func MainCall9() {
	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println("Func Main Call 9 , ", sum)
}

// 总结&分析
// Go是强数据类型，不同的类型运算，需要进行类型转化

// 把一个变量赋值给另一个不同类型的变量，需要显式的类型转换
func MainCall10() {
	i := 10
	var j float64 = float64(i) // 若没有显式转换，该语句会报错
	fmt.Printf("Func Main Call 10 , i 的值 %d, 类型是 %T \n", i, i)
	fmt.Printf("Func Main Call 10 , j 的值 %f, 类型是 %T \n", j, j)
}

// 总结&分析
// 运行结果如下
// Func Main Call 10 , i 的值 10, 类型是 int
// Func Main Call 10 , j 的值 10.000000, 类型是 float64
