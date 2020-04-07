package ch29

import (
	"errors"
	"fmt"
)

func MainCall10() {
	f101() //<nil>
	f102() //defer error
	f103() //<nil>
}

func f101() {
	var err error

	defer fmt.Println(err) // 参数传递

	err = errors.New("defer error")
	return
}

func f102() {
	var err error

	defer func() {
		fmt.Println(err)
	}()

	err = errors.New("defer error")
	return
}

func f103() {
	var err error

	defer func(err error) {
		fmt.Println(err)
	}(err)

	err = errors.New("defer error")
	return
}
