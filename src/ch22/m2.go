package ch22

import (
	"fmt"
	"time"
)

func Demo5(done chan bool)  {

	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
