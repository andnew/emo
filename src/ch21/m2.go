package ch21

import (
	"log"
	"time"
)

func Number() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Microsecond)
		log.Println("Number is ", i)
	}
}

func Alphabets() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Microsecond)
		log.Printf("Alphabets is %c\n", i)
	}
}
