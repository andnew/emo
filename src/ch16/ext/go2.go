package ext

import "fmt"

func MainExt2() {
	// 比较2个结构体
	demo01()
	// 不同类型的结构体变量是不能比较的：

}

type Employee6 struct {
	FirstName, LastName string
	salary              int
	fullTime            bool
}

//相同类型的结构比较
func demo01() {
	ross := Employee6{
		FirstName: "Ross",
		LastName:  "Bingo",
	}
	jack := Employee6{
		FirstName: "Jack",
		LastName:  "Lee",
	}
	fmt.Println(ross == jack)
}

// 总结&分析
// 如果结构体的所有成员都是可比较的，则这个结构体就是可比较的。可以使用 ==和 !=作比较，其中==是按照顺序比较两个结构体变量的成员变量

//不同类型的结构体变量是不能比较的：
type User struct {
	username string
}
type Employee7 struct {
	FirstName, LastName string
	salary              int
	fullTime            bool
}

func demo02() {
	ross := Employee7{
		FirstName: "Ross",
		LastName:  "Bingo",
	}
	user := User{
		username: "Seekload",
	}
	fmt.Println(ross, user)
	//fmt.Println(ross == user) //编译报错 Invalid operation: ross == user (mismatched types Employee7 and User)
}
