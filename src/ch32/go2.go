package ch32

import "fmt"

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName1(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func MainCall2() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName1(&firstName, nil)
	fmt.Println("returned normally from main")
}

// 总结&分析
// 函数 recover() 收集 panic函数发生的错误，从而终止 panic 续发事件。
