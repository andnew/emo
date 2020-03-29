package ext

import "fmt"

func MainSlice1() {
	//1、正确的声明
	makeOk()
	//2、使用字面量创建切片
	makeDirect()
	//3、只初始化某一个索引的值
	makeWithIndex()
	//4、基于已有的数组或者切片创建切片
	makeSliceWithArray()
	//5、示例1
	exam1()
}

func makeOk() {
	// 声明一个长度为3、容量为5的整型切片
	s1 := make([]int, 3, 5)
	fmt.Println("s1的 长度:", len(s1), ",容量:", cap(s1))

	//声明一个长度和容量都是5的字符串切片
	s2 := make([]string, 5)
	fmt.Println("s2的 长度:", len(s2), ",容量:", cap(s2))

	// len 必须 小于等于 cap ，否则编译报错  len larger than cap in make([]int)
	//s3 := make([]int , 5, 3)
	//fmt.Println(s3)
}

// 总结&分析
// 切片的声明是通过Go内置 make 声明，语法: s := make([]T, len, cap),其中 len <= cap

// 使用字面量创建切片
func makeDirect() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("使用字面量创建切片:", s, ", 长度:", len(s), ",容量:", cap(s))
	//使用字面量创建切片: [1 2 3 4 5] , 长度: 5 ,容量: 5

	s1 := [10]int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	fmt.Printf("s1的类型:%T,%v,长度:%d,容量:%d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("s2的类型:%T,%v,长度:%d,容量:%d\n", s2, s2, len(s2), cap(s2))
	//s1的类型:[10]int,[1 2 3 4 5 0 0 0 0 0],长度:10,容量:10
	//s2的类型:[]int,[1 2 3 4 5],长度:5,容量:5
}

//总结&分析
// 字面量创建切片的方式与创建数组的方式类似,只不过不用指定[]的值，这时候创建的切片长度和容量是相等的。

// 只初始化某一个索引的值
func makeWithIndex() {
	s := []int{4: 1}
	fmt.Println("只初始化某一个索引的值,s 数据是:", s, " 长度:", len(s), ",容量:", cap(s))
	//只初始化某一个索引的值,s 数据是: [0 0 0 0 1]  长度: 5 ,容量: 5

	s1 := []int{4: 1, 9: 9}
	fmt.Println("只初始化某一个索引的值,s1 数据是:", s1, " 长度:", len(s1), ",容量:", cap(s1))
	//只初始化某一个索引的值,s1 数据是: [0 0 0 0 1 0 0 0 0 9]  长度: 10 ,容量: 10
}

// 总结&分析
// 指定了第5个元素为1，其他元素初始化为 0

//基于已有的数组或者切片创建切片
func makeSliceWithArray() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("s[:]", s[:])
	fmt.Println("s[2:]", s[2:])
	fmt.Println("s[:4]", s[:4])
	fmt.Println("s[2:4]", s[2:4])
	//s[:] [0 1 2 3 4 5 6 7 8 9]
	//s[2:] [2 3 4 5 6 7 8 9]
	//s[:4] [0 1 2 3]
	//s[2:4] [2 3]

	s1 := s[2:4]
	fmt.Println(len(s1), cap(s1)) // 输出：2 8

	//使用3个索引的方法，第3个用来限定新切片的容量，用法为slice[i:j:k]。
	s2 := s[2:4:8]
	fmt.Println(s2, len(s2), cap(s2)) // 输出：[2 3]
}

//总结&分析
// 切片的操作符,[start, end] ，切片的索引是 0 ,所以 end的元素是 end - 1
// 对底层数组大小为k的切片执行[i,j]操作之后获得的新切片的长度和容量是：长度：j-i容量：k-i
// 官方总结
// 使用操作符[start,end]，简写成[i,j]，表示从索引i，到索引j结束，截取已有数组或者切片的任意部分，返回一个新的切片，新切片的值包含原切片的i索引的值，
// 但是不包含j索引的值。i、j都是可选的，i如果省略，默认是0，j如果省略，默认是原切片或数组的长度。i、j都不能超过这个长度值。

// 对底层数组大小为k的切片执行[i,j]操作之后获得的新切片的长度和容量是：长度：j-i容量：k-i
// 就拿上一个例子的s[2:4]来说，原切片底层数组大小是10，所以新切片的长度是4-2=2，容量是10-2=8。

// 长度和容量如何计算：长度j-i，容量k-i。所以切片s2的长度和容量分别是2、6。
// 注意：k不能超过原切片（数组）的长度，否则报错panic: runtime error: slice bounds out of range。例如上面的例子中，第三个索引值不能超过10

// 示例1
func exam1() {
	s := []int{0, 1, 2, 3, 4, 5}

	fmt.Println("before,s:", s)
	s1 := s[1:4]
	fmt.Println("before,s1:", s1)
	s1[1] = 10
	fmt.Println("after,s1:", s1)
	fmt.Println("after,s:", s)
}
