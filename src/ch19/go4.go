package ch19

import "fmt"

// 接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。
func MainCall4() {
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
		d1.Describe() //报错
	}
}

// 总结&分析
// 对于值为 nil 的接口，由于没有底层值和具体类型，当我们试图调用它的方法时，程序会产生 panic 异常。
// 执行结果
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109bb8f]
