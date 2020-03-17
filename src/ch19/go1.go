package ch19

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.Name, p.Age)
}

type Address struct {
	State, Country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s Country %s\n", a.State, a.Country)
}

func MainCall1() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()

	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}
	//d2 = a //Cannot use 'a' (type Address) as type Describer in assignment Type does not implement 'Describer' as 'Describe' method has a pointer receiver
	d2 = &a
	d2.Describe()
}

// 总结&分析
// 在讨论 方法 的时候就已经提到过，使用值接受者声明的方法，既可以用值来调用，也能用指针调用。
// 不管是一个值，还是一个可以解引用的指针，调用这样的方法都是合法的。
// d2 = a 报错的解释 对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。但接口中存储的具体值（Concrete Value）并不能取到地址，因此在第 45 行，对于编译器无法自动获取 a 的地址，于是程序报错。
