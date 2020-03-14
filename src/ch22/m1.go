package ch22

import "log"

// 信道的定义
func Demo1() {

	// 第一种方式进行声明
	var a chan int
	if a == nil {
		log.Println("channel a is nil, going to define")
		//使用make方式声明,示例化
		a = make(chan int)
		log.Printf("Type of is %T\n", a)
	}

	// 第二种方式声明
	b := make(chan int)
	log.Printf("b type is %T\n", b)
}

// 如何使用
// 通过信道进行发送和接收
func Demo2() {
	a := make(chan int)
	// 语法通过信道发送和接收数据
	data := <-a //读取信道 a
	a <- data   //写入信道 a
	//信道旁的箭头方向指定了是发送数据还是接收数据
	//在第二行，箭头对于 a 来说是向外指的，因此我们读取了信道 a 的值，并把该值存储到变量 data
	//在第三行，箭头指向了a，因此我们在把数据写入信道 a
}

// 信道的特征
func Demo3() {
	//发送与接收默认是阻塞的。这是什么意思？
	//当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
	//与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
	//
	//信道的这种特性能够帮助 Go 协程之间进行高效的通信，不需要用到其他编程语言常见的显式锁或条件变量。
	log.Println("Hello world goroutine.....")
}

// 使用channel 重写Demo3的代码
func Demo4(done chan bool) {
	log.Println("hello world goroutine, demo4 ....")
	// 如果没有数据写入，程序发生阻塞
	done <- true
	// 注释 done <- true 后的总结
	// 代码的错误提示
	//fatal error: all goroutines are asleep - deadlock!
	//
	//	goroutine 1 [chan receive]:
	//main.main()
	///Users/wangyong/workdirs/github/goes/emo/src/main.go:28 +0x111

}
