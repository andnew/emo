package ch15

import "fmt"

var ap *int

func MainCall1() {
	demo11()
}

func demo11() {

	fmt.Println("ap===", ap)
	a := 1 // define int
	b := 2 // define int
	fmt.Println("a=[", a, "],b=[", b, "]")
	ap = &a
	fmt.Println("ap=&a", ap)
	// set ap to address of a (&a)
	//   ap address: 0x2101f1018
	//   ap value  : 1

	*ap = 3
	fmt.Println("*ap = 3 ", ap, ", a =", a)
	// change the value at address &a to 3
	//   ap address: 0x2101f1018
	//   ap value  : 3

	a = 4
	fmt.Println("a = 4后的值 ", a, &a)
	// change the value of a to 4
	//   ap address: 0x2101f1018
	//   ap value  : 4
	ap = &b
	fmt.Println("ap = &b ,", ap, *ap)
	// set ap to the address of b (&b)
	//   ap address: 0x2101f1020
	//   ap value  : 2

	ap2 := ap
	fmt.Println("ap2=", ap2, *ap2, ap, *ap)

	*ap = 5
	fmt.Println("ap2=", ap2, *ap2, ap, *ap)

	ap = &a
	fmt.Println("ap2=", ap2, *ap2, ap, *ap, a, &a)
}

// 总结&分析
// 指针和引用 * 既是定义指针，又是解引用
// & 表示获取指针的引用
