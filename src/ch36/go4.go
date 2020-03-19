package ch36

import (
	"fmt"
	"os"
)

// 追加到文件

func MainCall4() {
	f, err := os.OpenFile("line", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}

// 总结&分析
// 文件的操作 os.O_APPEND|os.O_WRONLY,
// 文件的权限说明
