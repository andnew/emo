package ch36

import (
	"fmt"
	"os"
)

// 将字节写入文件
// 将字节写入文件和写入字符串非常的类似

func MainCall2() {
	f, err := os.Create("/Users/wangyong/xcript/bytes")
	if err != nil {
		fmt.Println(err)
		return
	}

	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 总结&分析
// 函数 Write 写入文件
