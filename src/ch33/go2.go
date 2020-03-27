package ch33

import "fmt"

func MainCall2() {
	func() {
		fmt.Println("hello world first class function")
	}()
}
