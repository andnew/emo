package ch32

import "fmt"

// 需要注意的是，你应该尽可能地使用错误，而不是使用 panic 和 recover。只有当程序不能继续运行的时候，才应该使用 panic 和 recover 机制。
// panic 有两个合理的用例。
// 发生了一个不能恢复的错误，此时程序不能继续运行。 一个例子就是 web 服务器无法绑定所要求的端口。在这种情况下，就应该使用 panic，因为如果不能绑定端口，啥也做不了。
// 发生了一个编程上的错误。 假如我们有一个接收指针参数的方法，而其他人使用 nil 作为参数调用了它。在这种情况下，我们可以使用 panic，因为这是一个编程错误：用 nil 参数调用了一个只能接收合法指针的方法。

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}

	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullname")
}

func MainCall1() {
	defer fmt.Println("deferred call in MainCall7")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("return normally from MainCall7")
}

// 总结&分析
// panic 是内部函数，将错误的信息提供，而终止程序的运行；
// 理解一下panic是如何工作的
// 程序发生 panic。当出现了 panic 时，程序就会终止运行，打印出传入 panic 的参数，接着打印出堆栈跟踪。
// 在发生panic地方，打印出日志信息，堆栈信息依次的向顶层抛出，打印调用的链信息

// 总结一下 panic 做了什么。
// 当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
// 这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪，最后程序终止。
