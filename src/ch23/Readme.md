package main

import (
	"ch23"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello emo!")

	// 第一个示例
	//ch23.Demo1()

	// 第二个示例
	//ch := make(chan int, 2)
	//go ch23.Demo2(ch)
	//
	//time.Sleep(2 * time.Second)
	//for v := range ch{
	//	fmt.Println("read value", v,"from ch")
	//	time.Sleep(2 * time.Second)
	//}

	// 第三个示例 死锁
	//ch23.Demo3()

	// 第四个示例 长度和容量
	//ch23.Demo4()

	// 第五个示例 WaitGroup
	//no := 3
	//var wg sync.WaitGroup
	//for i := 0; i < no; i++ {
	//	wg.Add(1)
	//	go ch23.Demo5(i, &wg)
	//}
	//wg.Wait()
	//fmt.Println("All go routines finished executing")
	// 总结
	// 上述程序里，for 循环迭代了 3 次(for i := 0; i < no; i++ { )，
	// 在循环内调用了 wg.Add(1) 。因此计数器变为 3。
	// for 循环同样创建了 3 个 process 协程[ go ch23.Demo5(i, &wg)]，
	// 然后在 代码行 调用了 wg.Wait()，确保 Go 主协程等待计数器变为 0。
	// 在 ch23.Demo5  行，Demo5 协程内调用了 wg.Done，可以让计数器递减。
	// 一旦 3 个子协程都执行完毕（即 wg.Done() 调用了 3 次），
	// 那么计数器就变为 0，于是主协程会解除阻塞。
	//
	//在 go ch23.Demo5(i, &wg) 代码行里，传递 wg 的地址是很重要的。
	//如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，因而当它们执行结束时，main 函数并不会知道。

	// 第六个示例 工作池
	startTime := time.Now()
	noOfJobs := 100
	go ch23.Allocate(noOfJobs)
	done := make(chan bool)
	go ch23.ResultFunc(done)
	noOfWorkers := 30
	ch23.CreateWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
