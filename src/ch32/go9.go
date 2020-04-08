package ch32

import (
	"errors"
	"fmt"
)

func MainCall9() {
	err := CallMe()
	if err != nil {
		fmt.Println(err.Error(), err)
	}
}

// 大写 C 开头，被外部调用
func CallMe() error {
	err := deal()
	if err != nil {
		errors.New("Some msg")
	}
	return err
}

// 小写 d 开头，只能被内部使用
func deal() (e *MyError) {
	if true {
		e = new(MyError)
		return
	}
	return
}
