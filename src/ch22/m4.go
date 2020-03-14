package ch22

import "fmt"

//信道 双向信道，即通过信道既能发送数据,又能接收数据。
// 其实也可以创建单向信道。这种信道只能发送或接收数据

func SendData(sendch chan<- int) {
	sendch <- 100
	fmt.Println("sendch====", sendch)
}
