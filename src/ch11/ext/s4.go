package ext

import "fmt"

func MainSlice4() {
	// nil 切片
	exam01()
	// 空切片
	exam02()
	// copy
	exam03()
	//
	exam04()
}

// nil 切片与空切片

func exam01() {
	// nil 切片
	var s []int
	fmt.Println(s == nil)
	fmt.Println(len(s), cap(s))
}

// 切片的零值是nil
// 切片的类型在初始化时已经确认，就是[]Type，上面的代码就声明了[]int类型的nil切片s。nil切片的指向底层数组的指针为nil。

// 空切片
// 声明空切片的方式有2种
func exam02() {
	//1、使用 make 创建空的整型切片
	s0 := make([]int, 0)
	s1 := []int{}
	fmt.Printf("输出s0 %v , 长度: %d , 容量: %d\n", s0, len(s0), cap(s0)) //输出s0 [] , 长度: 0 , 容量: 0
	fmt.Printf("输出s1 %v , 长度: %d , 容量: %d\n", s1, len(s1), cap(s1)) //输出s1 [] , 长度: 0 , 容量: 0
}

// 总结&分析
// 不管是使用 nil 切片还是空切片，对其调用内置函数append、len和cap的效果都是一样的。

//copy函数
// Go提供了内置函数copy，可以讲一个切片复制到另一个切片。函数原型：
// func copy(dst, src []Type) int
// dst是目标切片，src是源切片，函数返回两者长度的最小值。

func exam03() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	s4 := []int{1, 2, 3}
	// 1、
	n1 := copy(s1, s2) //因为s1是nil切片，执行完copy操作之后，s1依然还是nil
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("s1 == nil", s1 == nil)
	// 2、
	n2 := copy(s2, s3)
	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)
	// 3、
	n3 := copy(s3, s4)
	fmt.Printf("n3=%d, s3=%v, s4=%v\n", n3, s3, s4)
}

func exam04() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
