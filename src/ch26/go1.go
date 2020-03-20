package ch26

import "ch26/employee"

// Go 不支持类，而是提供了结构体。结构体中可以添加方法。这样可以将数据和操作数据的方法绑定在一起，实现与类相似的效果。

func MainCall1() {
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()
}

func MainCall2() {
	var e employee.Employee
	e.LeavesRemaining()
}

func MainCall3() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}

// 总结&分析
// Go 并不支持构造器。如
// 果某类型的零值不可用，需要程序员来隐藏该类型，避免从其他包直接访问。
// 程序员应该提供一种名为 NewT(parameters) 的函数，按照要求来初始化 T 类型的变量。
// 按照 Go 的惯例，应该把创建 T 类型变量的函数命名为 NewT(parameters)。这就类似于构造器了。
// 如果一个包只含有一种类型，按照 Go 的惯例，应该把函数命名为 New(parameters)， 而不是 NewT(parameters)。
