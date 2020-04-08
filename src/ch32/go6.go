package ch32

import (
	"fmt"
	"time"
)

func MainCall6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f61", r)
		}
	}()
	fmt.Println("MainCall7 start")
	go f61(2)
	// 当前 goroutine 睡眠，等待 panic
	time.Sleep(time.Second)
	fmt.Println("MainCall7 end")
	// 执行结果
	//MainCall7 start
	//Calling g61.
	//Printing in g61 2
	//Printing in g61 3
	//Panicking
	//Defer in g61 3
	//Defer in g61 2
	//panic: 4

	//goroutine 6 [running]:
}

func f61(i int) {
	fmt.Println("Calling g61.")
	g61(i)
	fmt.Println("Returned normally from g61.")
}

func g61(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g61", i)
	fmt.Println("Printing in g61", i)
	g61(i + 1)
}

//总结&分析
// 首先从程序的运行效果看，main方法没有执行完，而被终止。
// panic 并没有被解决，一直抛到最上层，使整个程序崩溃，
// panic 只能由当前 goroutine的 recover 解决。
