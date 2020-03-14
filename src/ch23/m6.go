package ch23

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 缓冲信道的重要应用之一就是实现工作池
// 一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配

//工作池的核心功能如下:
// 1、创建一个Go协程池，监听一个等待作业分配的输入型缓冲信道
// 2、将作业添加到该输入型缓冲信道中
// 3、作业完成后，再将结果写入一个输出型缓冲信道
// 4、从输出型缓冲信道读取并打印结果。

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

// 创建 用于接收作业和写入结果的缓冲信道
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 每一个整数的每一位之和
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 创建工作协程的函数
// 函数创建了一个工作者(worker)，读取 jobs 信道的数据，根据当前的 jobs 和 digits 函数的返回值，
// 创建了一个 Result 结构体变量，然后将结果写入 results 缓冲信道。 worker 函数接收了一个 WaitGroup 类型的 wg 作为参数，
// 当所有的 jobs 完成的时候，调用了 Done() 方法。
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// 函数创建一个Go协程的工作池
// 函数参数是需要创建的工作协程的数量。在创建Go协程之前，它调用了 wg.Add(1)方法，于是 WaitGroup 计数器递增。
// 接下来,创建工作协程，并向 worker 函数传递 wg的地址。创建了需要的工作协程后，函数调用 wg.Wait(),等待所有的Go协程执行完毕。
// 所有协程完成之后，函数关闭 results 信道。因为所有协程已经都已经执行完毕，于是不再需要向 results 信道写入数据了。
func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// 把任务分批给工作者
// 函数接收所需要创建的作业数量作为输入参数，生成了最大值为998的伪随机数，并使用该随机数创建了 Job 结构体变量。
// 这个函数把 for 循环 的计数器 i 作为 id ， 最后把创建的结构体变量写入 jobs 信道。 当写入所有的 job 时，它关闭 jobs 信道。
func Allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 读取 results 信道，并打印 jobs 的 id、 输入的随机数的每位数之和。
// result 函数也接收 done 信道作为参数，当打印所有结果时，done 会被写入 true
func ResultFunc(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
