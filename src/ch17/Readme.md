package main

import (
	"ch17"
	"config"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello emo!")
	//config.LoadConfig()
	config.LoadConfig()

	// 示例1 说明方法定于&使用

	// 正常的构造器参数
	emp1 := ch17.Employee{Name: "Jame", Salary: 1000, Currency: "$"}
	// 构造器参数顺序发生变化
	emp2 := ch17.Employee{Salary: 2000, Currency: "$", Name: "Jame2"}
	// 通过属性函数进行赋值
	emp3 := ch17.Employee{}
	emp3.Name = "Sam Adolf"
	emp3.Currency = "$"
	emp3.Salary = 5000
	emp1.DisplaySalary()
	emp2.DisplaySalary()
	emp3.DisplaySalary()

	// 示例2  方法与函数的区别
	// Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径
	// 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
	// 假设我们一个Rectangle和Circle结构体。可以在Rectangle和Circle上分别定义一个Area方法。
	r := ch17.Rectangle{Length: 10, Width: 5}
	log.Printf("Area of rectangle %d\n", r.Area())
	c := ch17.Circle{Radius: 12}
	log.Printf("Area of circle %f\n", c.Area())

	// 指针接收器与值接收器
	// 示例1 和 示例2 使用的是值接收器
	// 值接收器和指针接收器之间的区别是 ， 在指针接收器的方法内部的改变对于调用者是可见的。
	// 然而值接收器的情况不是这样的。
	// 示例3 指针接收器和值接收器的区别

	e := ch17.Employee1{Name:"Mark Andrew",Age:100}
	log.Printf("示例3 指针接收器和值接收器的区别\n")
	log.Printf( "Employee name before change: %s", e.Name)
	e.ChangeName("Michael Andrew")
	log.Printf("Employee name after change: %s", e.Name)

	log.Printf("Employee age before change: %d", e.Age)
	//(&e).ChangeAge(51)
	e.ChangeAge(51) // Go语言自动解析为 (&e).ChangeAge(51)
	log.Printf("Employee age after change: %d", e.Age)

	// 总结
	// 什么时候使用指针接收器，什么时候使用值接收器?
	// 一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。
	// 指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。
	// 在其他的所有情况，值接收器都可以被使用。
}
