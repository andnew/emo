package main

import (
	"ch22"
	"config"
	"fmt"
)

func main() {
	fmt.Println("Hello emo!")
	//config.LoadConfig()
	config.LoadConfig()

	// 第一个示例
	//ch22.Demo1()
	// channel的使用
	//ch22.Demo2()
	// 第三个示例
	//使用channel前的代码处理方式
	//go ch22.Demo3()
	//time.Sleep(1 * time.Second)
	//log.Println("main function....")

	// 第四个示例
	//done := make(chan  bool)
	//go ch22.Demo4(done)
	//<- done
	//log.Println("main function over ..")
	// 总结 代码种 创建了一个bool类型的信道 done, 并把 done 作为参数传递给了 Demo4 协程。
	//  <- done 代码行， 通过信道 done 接收数据。这行代码发生了阻塞，除非有协程向done写入数据,否则程序不会跳到下一行代码。
	// 这就不需要用以前的 time.Sleep 来阻止 Go 主协程退出了。
	//现在我们的 Go 主协程发生了阻塞，等待信道 done 发送的数据。该信道作为参数传递给了协程 Demo4，
	//Demo4 打印出 Hello world goroutine，接下来向 done 写入数据。当完成写入时，Go 主协程会通过信道 done 接收数据，
	//于是它解除阻塞状态，打印出文本 main function。

	//第五个示例
	//done := make(chan bool)
	//fmt.Println("Main going to call Demo5 go goroutine")
	//go ch22.Demo5(done)
	//<-done
	//fmt.Println("Main received data")

	//第六个示例
	//number := 589
	//sqrch := make(chan int)
	//cubech := make(chan int)
	//go ch22.CalcSquares(number, sqrch)
	//go ch22.CalcCubes(number, cubech)
	//
	//squares, cubes := <-sqrch, <-cubech
	//fmt.Println("Final output", squares+cubes)
	// 总结
	//  函数 CalcSquares 计算一个数每位的平方和，并把结果发送给信道 squareop.
	//  同理 CalcCubes 计算一个数每位的立方和，并把结果发送给信道 cubeop
	//  这两个函数分别在单独的协程里运行，每个函数都有传递信道的参数，以便写入数据。
	//  Go主协程会在 squares, cubes := <-sqrch, <-cubech 两个信道传来的数据。
	//  一旦从两个信道接收完数据，数据就会存储在变量 squares 和 cubes 里，然后计算并打印出最后结果。

	// 第七个示例
	// 死锁
	//  当Go协程给一个信道发送数据时，会有其他Go协程来接收数据。如果没有的话，程序就会在运行时触发panic，形成死锁。
	//  同理，当有Go协程等着从一个信道接收数据时，我们期望其他的Go协程会向该信道写入数据，要不然程序就会触发panic
	//ch := make(chan int)
	//ch <- 5
	// 总结 我们创建了一个信道 ch ,接收在下一行 ch <- 5,把 5 发送到这个信道。 对于 代码
	// ch := make(chan int)
	// ch <- 5
	// 没有其他的协程从 ch 接收数据。于是程序触发了 panic, 出现如下错误。
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send]:
	// main.main()
	// /Users/wangyong/workdirs/github/goes/emo/src/main.go:64 +0x3e5

	// 第八个示例
	// 单向信道
	//sendch := make(chan <- int)
	//go ch22.SendData(sendch)
	//fmt.Println(<-sendch)
	//./main.go:78:14: invalid operation: <-sendch (receive from send-only type chan<- int)
	//总结

	// 第九个示例
	//cha1 := make(chan int)
	//fmt.Println("cha1====",cha1)
	//go ch22.SendData(cha1)
	//fmt.Println(<- cha1)

	// 第十个示例
	//ch := make(chan int)
	//go ch22.Producter(ch)
	//for {
	//	v, ok := <-ch
	//	if ok == false {
	//		break
	//	}
	//	fmt.Println("Received ", v, ok)
	//}
	//总结
	// v, ok := <-ch 如果成功接收信道所发送的数据，那么ok 等于 true。
	// 而如果 ok 等于 false，说明我们试图读取一个关闭的通道。
	// 从关闭的信道读取到的值会是该信道类型的零值

	// 第十一个示例
	//ch := make(chan int)
	//go ch22.Producter(ch)
	//for v := range ch {
	//	fmt.Println("Received ", v)
	//}
	// 总结
	// for range 循环从信道 ch 接收数据，直到该信道关闭。一旦关闭了 ch，循环会自动结束。该程序会输出

	// 用信道,for range方式 重写 第六示例
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go ch22.CalcSquares2(number, sqrch)
	go ch22.CalcCubes2(number, cubech)

	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
