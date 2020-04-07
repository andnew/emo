package ch29

import "fmt"

func MainCall8() {
	f1 := f1()
	fmt.Println("f1()====", f1)
	f2 := f2()
	fmt.Println("f2()====", f2)
	f3 := f3()
	fmt.Println("f3()====", f3)
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0 //1
}

// f1 的解读
//func f1() (r int) {
//
//	// 1.赋值
//	r = 0
//
//	// 2.闭包引用，返回值被修改
//	defer func() {
//		r++
//	}()
//
//	// 3.空的 return
//	return
//}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t // 5
}

// f2 的解读
//func f2() (r int) {
//    t := 5
//    // 1.赋值
//    r = t
//
//    // 2.闭包引用，但是没有修改返回值 r
//    defer func() {
//        t = t + 5
//    }()
//
//    // 3.空的 return
//    return
//}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1 // 1
}

//f3 的解读
//func f3() (r int) {
//
//    // 1.赋值
//    r = 1
//
//    // 2.r 作为函数参数，不会修改要返回的那个 r 值
//    defer func(r int) {
//        r = r + 5
//    }(r)
//
//    // 3.空的 return
//    return
//}
