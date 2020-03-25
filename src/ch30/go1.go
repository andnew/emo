package ch30

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// 错误表示程序中出现了异常情况。
// 错误用内建的 error 类型来表示。

func MainCall1() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("MainCall1========", f.Name(), "opened successfully")
}

func MainCall2() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("MainCall2========", "File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func MainCall3() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

func MainCall4() {
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

func MainCall5() {
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files)
}

// 总结与分析
// 1、什么是错误？
// 2、错误的表示
// 3、获取错误详细信息的各种方法
// 4、不能忽视错误
