package ch32

import "fmt"

func MainCall5() {

	fmt.Println("MainCall7 start")
	f51(2)
	fmt.Println("MainCall7 end")
	// 执行结果
	//MainCall7 start
	//Calling g61.
	//Printing in g61 2
	//Printing in g61 3
	//Panicking
	//Defer in g61 3
	//Defer in g61 2
	//Recovered in f61 4
	//MainCall7 end
}

func f51(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f61", r)
		}
	}()
	fmt.Println("Calling g61.")
	g51(i)
	fmt.Println("Returned normally from g61.")
}

func g51(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g61", i)
	fmt.Println("Printing in g61", i)
	g51(i + 1)
}

//总结&分析
// defer 是一个先进后出的调用栈，defer 关键词后面跟的语句或函数。这些参数中存在2类，1、值传递的参数，2、是匿名参数
