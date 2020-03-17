package ch11

import "fmt"

func slice1() {
	aint := [3]int{1, 2, 3}
	aslice := aint[1:2]
	aslice = append(aslice, 4)
	aslice[0] = 99
	fmt.Println("slice1 == ", aint, aslice, len(aslice), cap(aslice))
}

func slice2() {
	aint := [3]int{1, 2, 3}
	aslice := aint[1:2]
	fmt.Println("首次===== ", &aslice, len(aslice), cap(aslice))
	aslice = append(aslice, 4, 5)
	fmt.Println("after append ", &aslice, len(aslice), cap(aslice))
	aslice[0] = 99
	fmt.Println("update aslice[0] ", &aslice, len(aslice), cap(aslice))
	fmt.Println("slice2 == ", aint, aslice, len(aslice), cap(aslice))
}
