package ext

import "fmt"

// Go 语言的数组可以存储一组相同类型的数据，而结构体可以将不同类型的变量数据组合在一起，每一个变量都是结构体的成员。
// 创建并初始化一个结构体
// 可以使用下面的语法创建一个结构体：
// type StructName struct{
//     field1 fieldType1
//     field2 fieldType2
// }

func MainExt1() {
	// 声明
	create()
	// 结构体指针
	pointer()
	// 匿名成员
	noNameStruct()
	// 结构体嵌套
	demo()
	// 简化变量
	demo2()
}

func pointer() {
	rossPointer := &Employee{
		firstName: "ross",
		lastName:  "Bingo",
		salary:    1000,
		fullTime:  true,
	}
	fmt.Println(rossPointer, *rossPointer)
	fmt.Println("firstName:", (*rossPointer).firstName)
	fmt.Println("firstName:", rossPointer.lastName)
}

//总结&分析
// ross_pointer是一个结构体变量，所以(*ross_pointer).firstName和ross_pointer.lastName都是正确的访问方式 。

// 匿名成员
// 定义结构体时可以只指定成员类型，不用指定成员名，Go 会自动地将成员类型作为成员名。这种结构体成员称为匿名成员。这个结构体成员的类型必须是命名类型或者是指向命名类型的指针。
type Week struct {
	string
	int
	bool
}

func noNameStruct() {
	week := Week{"Friday", 1000, true}
	fmt.Println(week)

	e := Employee1{
		firstName: "Hello",
		lastName:  "World",
		salary:    0,
		bool:      false,
	}
	fmt.Println(e)
}

func create() {

	fmt.Println("Ext 1 Code....")
	// 使用变量/别名的方式赋值
	var ross Employee
	ross.firstName = "ross"
	ross.lastName = "Bingo"
	ross.salary = 1000
	ross.fullTime = true
	fmt.Println(ross)

	// 还可以使用字面量的方式初始化结构体
	//1、方式一 按照结构体的字段顺序赋值
	// 初始化时省略了成员变量名称，但是必须按顺序地将给出所有的成员的值。必须记住所有成员的类型且按顺序赋值
	ross1 := Employee{
		"ross",
		"Bingo",
		1000,
		true,
	}
	fmt.Println(ross1)
	//输出：{ross Bingo 1000 true}

	//2、方式二 使用字段指定方式赋值
	// 不用关心成员变量的顺序，给需要初始化的成员赋值，未赋值的成员默认就是类型对应的零值
	ross2 := Employee{
		lastName:  "Bingo",
		firstName: "ross",
		salary:    1000,
	}
	fmt.Println(ross2)
}

// 总结&分析
// 方式1 和 方式2 不可以混搭
// 成员变量的顺序对于结构体的同一性很重要，如果将上面的 firstName、lastName 互换顺序或者将 fullTime、salary 互换顺序，都是在定义一个不同的结构体类型

type Employee struct {
	firstName string // firstName,lastName 数据类型一致可以放在一行上,代码 firstName,lastName string
	lastName  string
	salary    int
	fullTime  bool
}

type Employee1 struct {
	firstName, lastName string
	salary              int
	bool
}

type Salary1 struct {
	basic      int
	workerTime int
}

type Employee3 struct {
	firstName, lastName string
	salary              Salary1
	bool
}

func demo() {
	ross := Employee3{
		firstName: "Ross",
		lastName:  "Bingo",
		bool:      true,
		salary:    Salary1{1000, 100},
	}
	fmt.Println(ross.salary.basic)
}

// 采用匿名成员方式重新定义结构体类型 Employee3
type Employee4 struct {
	firstName, lastName string
	Salary1             // 匿名成员
	bool
}

func demo2() {
	ross := Employee4{
		firstName: "Ross",
		lastName:  "Bingo",
		bool:      true,
		Salary1:   Salary1{1000, 100},
	}
	fmt.Println(ross.basic)         // 访问方式一
	fmt.Println(ross.Salary1.basic) // 访问方式二
	ross.basic = 1200
	fmt.Println(ross.basic) // update
}

type Employee5 struct {
	firstName, lastName string
	Salary1
	basic int
	bool
}

func demo3() {
	ross := Employee5{
		firstName: "Ross",
		lastName:  "Bingo",
		bool:      true,
		basic:     200,
		//Salary1{1000, 100},
	}
	fmt.Println(ross.basic)
}

//我们修改了结构体 Employee5，多添加了一个与 Salary.basic同名的成员，但是编译出错：mixture of field:value and value initializers
