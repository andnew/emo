package ext

import "fmt"

//删除切片中的元素
func MainSlice6() {

	s := []int{1, 2, 3, 4, 5, 6}
	s = append(s[:2], s[3:]...) // 删除索引为2的元素
	fmt.Println(s)
}

//总结&分析
// Go没有提供删除切片元素的函数,采用变相获取删除的
// 执行结果是 [1 2 4 5 6]
