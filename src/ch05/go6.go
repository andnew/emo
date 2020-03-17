package ch05

import "fmt"

// 数字常量可以在表达式中自由混合和匹配，只有当它们被分配给变量或者在需要类型的代码中的任何地方使用时，才需要类型。

func MainCall9() {
	var a = 5.9 / 8
	fmt.Printf("a's type %T value %v", a, a)
}
