package ch32

import "fmt"

func MainCall7() {
	fmt.Println("MainCall7 start")
	f71(2)
	fmt.Println("MainCall7 end")
	// 执行结果
	//MainCall7 start
	//Calling g71.
	//Printing in g71 2
	//Printing in g71 3
	//Panicking
	//Defer in g71 3
	//Defer in g71 2
	//panic: 4

	//goroutine 1 [running]:

}

func f71(i int) {
	fmt.Println("Calling g71.")
	g71(i)
	fmt.Println("Returned normally from g71.")
	r := recover()
	fmt.Println("Recovered in f71", r)
}

func g71(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g71", i)
	fmt.Println("Printing in g71", i)
	g71(i + 1)
}

// 总结&分析
// recover 只能在 defer 中有效 , 并且 语句的位置应该放在调用语句前面
// defer panic recover ,recover 必须放在defer中使用，同时panic 和 recover 必须放在同一个goroutine中
