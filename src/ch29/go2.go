package ch29

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}

func MainCall2() {
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}

// 总结&分析
// defer 第二个示例，p.fullName ,延迟， 代码的运行效果是
// Welcome John Smith
