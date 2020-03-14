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

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

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

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func Allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func ResultFunc(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
