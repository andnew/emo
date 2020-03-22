package ch29

import "fmt"

// defer 语句的用途是：含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数。这

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func MainCall1() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}

// 总结&分析
// 一般的程序都是从左到右，从上到下执行 ，
// 代码的执行效果是
// Started finding largest
// Largest number in [78 109 2 563 300] is 563
// Finished finding largest

// 从示例中查看 , defer 的就是延迟执行
