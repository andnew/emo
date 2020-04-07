package ch29

import "fmt"

func MainCall9() {
	demo91()
	demo92()
}

func demo91() {
	var whatever [3]struct{}

	for i := range whatever {
		fmt.Println("循环====", i)
		defer func() {
			fmt.Println(i)
		}()
	}
	// 执行结果是
	//循环==== 0
	//循环==== 1
	//循环==== 2
	//2
	//2
	//2
}

type number int

func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }

func demo92() {
	var n number

	defer n.print()
	defer n.pprint()
	defer func() { n.print() }()
	defer func() { n.pprint() }()

	n = 3
	// 执行结果
	//3
	//3
	//3
	//0
}

// 总结&分析
// defer&闭包
