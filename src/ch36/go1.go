package ch36

import (
	"fmt"
	"os"
)

// 将字符串写入文件
// 最常见的写文件就是将字符串写入文件。这个写起来非常的简单。这个包含以下几个阶段。

// 创建文件
// 将字符串写入文件

func MainCall1() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
}

// 总结&分析
// 首先 os.Create(file) 创建的这个文件放在项目目录下
// os.Create的文件后，获取文件句柄, 文件的写入操作 WriteString 函数
//
