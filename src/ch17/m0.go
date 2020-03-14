package ch17

import "log"

// 方法的定义与使用

type Employee struct {
	Name     string
	Salary   int
	Currency string
}
// 方法的定义 func (t Type) methodName (parameter list){}
func (e Employee) DisplaySalary() {
	log.Printf("Salary of %s %s%d", e.Name, e.Currency, e.Salary)
}
