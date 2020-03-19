package ch36

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

// 并发写文件
// 当多个 goroutines 同时（并发）写文件时，我们会遇到竞争条件(race condition)。因此，当发生同步写的时候需要一个 channel 作为一致写入的条件。

// 我们将写一个程序，该程序创建 100 个 goroutinues。每个 goroutinue 将并发产生一个随机数，届时将有 100 个随机数产生。这些随机数将被写入到文件里面。我们将用下面的方法解决这个问题 .

// 创建一个 channel 用来读和写这个随机数。
// 创建 100 个生产者 goroutine。每个 goroutine 将产生随机数并将随机数写入到 channel 里。
// 创建一个消费者 goroutine 用来从 channel 读取随机数并将它写入文件。这样的话我们就只有一个 goroutinue 向文件中写数据，从而避免竞争条件。
// 一旦完成则关闭文件。

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func MainCall5() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
