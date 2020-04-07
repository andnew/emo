package ch29

import "fmt"

func MainCall11() {
	var a = Accumulator()

	fmt.Printf("%d\n", a(1))
	fmt.Printf("%d\n", a(10))
	fmt.Printf("%d\n", a(100))

	fmt.Println("------------------------")
	var b = Accumulator()

	fmt.Printf("%d\n", b(1))
	fmt.Printf("%d\n", b(10))
	fmt.Printf("%d\n", b(100))
	//执行结果
	//(0xc00001c098, 0) - 1
	//(0xc00001c098, 1) - 11
	//(0xc00001c098, 11) - 111
	//------------------------
	//(0xc00001c0d8, 0) - 1
	//(0xc00001c0d8, 1) - 11
	//(0xc00001c0d8, 11) - 111

}

func Accumulator() func(int) int {
	var x int

	return func(delta int) int {
		fmt.Printf("(%+v, %+v) - ", &x, x)
		x += delta
		return x
	}
}
