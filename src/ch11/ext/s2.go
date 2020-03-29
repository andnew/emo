package ext

import "fmt"

// 使用切片
func MainSlice2() {

	//示例1
	useExam1()
	//示例2
	//useExamWithEffect()
	//示例3
	appendExam()
	//示例4
	appendExam2()
	//示例5
	appendExam3()
}

// 切片的使用方法与数组的使用方法类似，直接通过索引就可以获取、修改元素的值。
//

func useExam1() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[1]) // 获取值   输出：2
	s[1] = 10         // 修改值
	fmt.Println(s)    //输出：[1 10 3 4 5]
}

// 只能访问切片长度范围内的元素，否则报错
func useExamWithEffect() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[2:3]
	fmt.Println(s1[1])
	//	panic: runtime error: index out of range [1] with length 1
}

// 使用append 追加一个示例
func appendExam() {
	s := []int{1, 2, 3, 4, 5}
	newS := s[2:4]
	newS = append(newS, 50)
	fmt.Println(s, newS)
	fmt.Println(&s[2] == &newS[0])
}

func appendExam2() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := s[2:4]
	fmt.Printf("before -> s=%v\n", s)
	fmt.Printf("before -> s1=%v\n", s1)
	fmt.Printf("before -> len=%d, cap=%d\n", len(s1), cap(s1))
	fmt.Println("&s[2] == &s1[0] is", &s[2] == &s1[0])

	s1 = append(s1, 60, 70, 80, 90, 100, 110)
	fmt.Printf("after -> s=%v\n", s)
	fmt.Printf("after -> s1=%v\n", s1)
	fmt.Printf("after -> len=%d, cap=%d\n", len(s1), cap(s1))
	fmt.Println("&s[2] == &s1[0] is", &s[2] == &s1[0])
}

// 总结&分析
// 1、从结果可以看出，切片的底层数组没有足够的可用容量， append函数会创建一个新的底层数组，将原数组的值复制到新数组里，再追加新的值，就不会影响原来的底层数组。
// 2、append函数会智能地增加底层数组的容量，目前的算法是：当数组容量<=1024时，会成倍地增加；当超过1024，增长因子变为1.25，也就是说每次会增加25%的容量。
// 3、注意
// 一般我们在创建新切片的时候，最好要让新切片的长度和容量一样，这样我们在追加操作的时候就会生成新的底层数组，和原有数组分离，就不会因为共用底层数组而引起奇怪问题，因为共用数组的时候修改内容，会影响多个切片。
//Go提供了...操作符，允许将一个切片追加到另一个切片上：

func appendExam3() {
	s := []int{1, 2, 3, 4, 5}
	s1 := []int{6, 7, 8}
	s = append(s, s1...)
	fmt.Println(s, s1)
}
