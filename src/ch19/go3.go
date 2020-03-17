package ch19

import "fmt"

// 接口的嵌套
// 尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。

type SalaryCalculator3 interface {
	DisplaySalary3()
}

type LeaveCalculator3 interface {
	CalculateLeavesLeft3() int
}

type EmployeeOperations interface {
	SalaryCalculator3
	LeaveCalculator3
}

type Employee3 struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee3) DisplaySalary3() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee3) CalculateLeavesLeft3() int {
	return e.totalLeaves - e.leavesTaken
}

func MainCall3() {
	e := Employee3{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary3()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft3())
}

// 总结&分析
// 接口包装
