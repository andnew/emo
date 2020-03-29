package ext

import "fmt"

func MainSlice3() {
	//使用for 循环迭代切片，配合 len 函数
	examFor1()
	//使用 for range 迭代切片
	examFor2()
	//注意示例
	examFor3()
}

func examFor1() {

	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++ {
		fmt.Printf("Index:%d,Value:%d\n", i, s[i])
	}
}

func examFor2() {
	fmt.Println("====examFor2====")
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Printf("Index:%d,Value:%d\n", i, v)
	}

	// 使用‘_’可以忽略返回值
	s1 := []int{1, 2, 3, 4, 5}
	for _, v := range s1 {
		fmt.Printf("Value:%d\n", v)
	}
}

// 需要注意的是，range返回的是切片元素的复制，而不是元素的引用。如果使用该值变量的地址作为指向每个元素的指针，就会造成错误。
func examFor3() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Printf("v:%d, v_addr:%p, elem_addr:%p,=== %d , %p\n", v, &v, &s[i], i, &i)
	}
}
