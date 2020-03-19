package ch36

import (
	"fmt"
	"os"
)

//将字符串一行一行的写入文件
//另外一个常用的操作就是将字符串一行一行的写入到文件。
func MainCall3() {
	f, err := os.Create("line")
	if err != nil {
		fmt.Println(err)
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.",
		"It is easy to learn Go."}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

// 总结&分析
// 用迭代并使用 for rang 循环这个数组，并使用 Fprintln Fprintln 函数 将 io.writer 做为参数，并且添加一个新的行，这个正是我们想要的
