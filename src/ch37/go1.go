package ch37

import "fmt"

func MainCall1() {
	// itoa 只能使用在表达式中
	demo01()
	// 每次 const 出现时，都会让 iota 初始化为0.【自增长】
	demo02()
	// 自定义常量
	demo03()
	// 可跳跃的值
	demo04()
	// 位掩码表达式
	demo05()
	// 定义数量级
	demo06()
	demo07()
	// 定义在一行
	demo08()
	// 中间插队
	demo09()
	// 常规使用
	demo10()
	// 验证示例
	demo11()
}

func demo01() {
	fmt.Println("iota只能使用在表达式中")
	//fmt.Println(iota) //编译报错 undefined: iota
}

// 总结&分析
// iota只能使用在表达式中

func demo02() {
	const a = iota // a=0
	const (
		b = iota //b=0
		c        //c=1
		d        //d= 2
	)

	fmt.Printf("a=%d\nb=%d\nc=%d\nd=%d\n", a, b, c, d)
}

// 总结&分析
// 每次 const 出现时，都会让 iota 初始化为0.【自增长】

// 自定义常量
func demo03() {
	//自增长常量经常包含一个自定义枚举类型，允许你依靠编译器完成自增设置。

	type Stereotype int

	const (
		TypicalNob            Stereotype = iota // 0
		TypicalHipster                          // 1
		TypicalUnixWizard                       // 2
		TypicalStartupFounder                   // 3
	)

	fmt.Printf("TypicalNob=%d\nTypicalHipster=%d\nTypicalUnixWizard=%d\nTypicalStartupFounder=%d\n", TypicalNob, TypicalHipster, TypicalUnixWizard, TypicalStartupFounder)
}

// 可跳过的值
func demo04() {
	//如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。
	type AudioOutput int

	const (
		OutMute   AudioOutput = iota // 0
		OutMono                      // 1
		OutStereo                    // 2
		_
		_
		OutSurround // 5
	)
	fmt.Println("OutSurround=", OutMute)
	fmt.Println("OutMono=", OutMono)
	fmt.Println("OutStereo=", OutStereo)
	fmt.Println("OutSurround=", OutSurround)
}

// 总结&分析
// 可跳过的值, 使用 _

// 位掩码表达式
func demo05() {
	fmt.Println("位掩码表达式")
	type Allergen int

	const (
		IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
		IgChocolate                         // 1 << 1 which is 00000010
		IgNuts                              // 1 << 2 which is 00000100
		IgStrawberries                      // 1 << 3 which is 00001000
		IgShellfish                         // 1 << 4 which is 00010000
	)
	//v := IgShellfish | IgEggs
	fmt.Printf("IgEggs=%b\n", IgEggs)
	fmt.Printf("IgChocolate=%b\n", IgChocolate)
	fmt.Printf("IgNuts=%b\n", IgNuts)
	fmt.Printf("IgStrawberries=%b\n", IgStrawberries)
	fmt.Printf("IgShellfish=%b\n", IgShellfish)
}

// 定义数量级
func demo06() {
	fmt.Println("定义数量级")
	type ByteSize float64

	const (
		_           = iota             // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
		EB                             // 1 << (10*6)
		ZB                             // 1 << (10*7)
		YB                             // 1 << (10*8)
	)
	fmt.Println("KB=", KB)
	fmt.Println("MB=", MB)
	fmt.Println("GB=", GB)
	fmt.Println("TB=", TB)
	fmt.Println("PB=", PB)
	fmt.Println("EB=", EB)
	fmt.Println("ZB=", ZB)
	fmt.Println("YB=", YB)
}

func demo07() {
	var data interface{} = 2
	fmt.Printf("%b\n", data)
	fmt.Printf("%d\n", data)
	fmt.Printf("%x\n", data)
	fmt.Printf("%f\n", data)
}

// 定义在一行的情况
func demo08() {
	const (
		Apple, Banana     = iota + 1, iota + 2
		Cherimoya, Durian // = iota + 1, iota + 2
		Elderberry, Fig   //= iota + 1, iota + 2
	)
	fmt.Println("Apple=", Apple)
	fmt.Println("Banana=", Banana)
	fmt.Println("Cherimoya=", Cherimoya)
	fmt.Println("Durian=", Durian)
	fmt.Println("Elderberry=", Elderberry)
	fmt.Println("Fig=", Fig)
}

// 总结&分析
// 定义在一行

// 中间插队
func demo09() {
	const (
		i = iota
		j = 3.14
		k = iota
		l
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	//那么打印出来的结果是 i=0,j=3.14,k=2,l=3
}

// 总结&分析

func demo10() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出
	)

	fmt.Println("Sunday=", Sunday)
	fmt.Println("Monday=", Monday)
	fmt.Println("Tuesday=", Tuesday)
	fmt.Println("Wednesday=", Wednesday)
	fmt.Println("Thursday=", Thursday)
	fmt.Println("Friday=", Friday)
	fmt.Println("Saturday=", Saturday)
	fmt.Println("numberOfDays=", numberOfDays)
}

func demo11() {

	const (
		h = iota //h=0
		i = 100  //i=100
		j        //j=100
		k = iota //k=3
	)
	fmt.Println("h=", h)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
}
