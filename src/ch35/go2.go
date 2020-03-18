package ch35

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func MainCall4() {

	fmt.Println("==========MailCall4============")

	f, err := os.Open("/Users/wangyong/workdirs/github/goes/emo/src/ch35/test.txt")

	if err != nil {
		log.Fatal("--------", err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// 方式1 按照行读取文件
	//br := bufio.NewReader(f)
	//for {
	//	a, _, c := br.ReadLine()
	//	if c == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(a))
	//}
	// 方式2 按照行读取文案
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// 当 Scan 返回 false 时，除非已经到达文件末尾（此时 Err() 返回 nil），否则 Err() 就会返回扫描过程中出现的错误。
// Scan() 方法读取文件的下一行，如果可以读取，就可以使用 Text() 方法。
