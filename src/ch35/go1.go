package ch35

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func MainCall1() {
	data, err := ioutil.ReadFile("./src/ch35/test.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

// 总结&分析
// 1、ioutil.ReadFile(filename),其中filename 方式1 要不就是绝对路径，方式2、使用 src下的相对路径
//  方式3、使用命令行标记传递文件路径

func MainCall2() {
	data, err := ioutil.ReadFile("/Users/wangyong/workdirs/github/goes/emo/src/ch35/test.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("==========MailCall2============\nContents of file:", string(data))
}

func MainCall3() {
	fmt.Println("===================main call 3================")
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	//flag.Parse()
	fmt.Println("value of fpath is", *fptr)
	data, err := ioutil.ReadFile(*fptr)

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
