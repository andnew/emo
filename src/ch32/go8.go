package ch32

import "fmt"

func MainCall8() {
	if err := try(); err != nil {
		fmt.Println("err != nil")
	}
}

type MyError struct{}

func (e *MyError) Error() string {
	return "Oooops"
}

func try() error {
	var err *MyError = nil
	if true {
		err = new(MyError)
	}
	return err
}
