package ch24

import (
	"fmt"
	"time"
)

// select 的应用
// 假设我们有一个关键性应用，需要尽快地把输出返回给用户。这个应用的数据库复制并且存储在世界各地的服务器上。
// 假设函数 server1 和 server2 与这样不同区域的两台服务器进行通信。每台服务器的负载和网络时延决定了它的响应时间。
// 我们向两台服务器发送请求，并使用 select 语句等待相应的信道发出响应。select 会选择首先响应的服务器，而忽略其它的响应。
// 使用这种方法，我们可以向多个服务器发送请求，并给用户返回最快的响应了。:）

func process(ch chan string) {
	time.Sleep(10500 * time.Microsecond)
	ch <- "Process successful"
}

func MainCall2() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Microsecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received, ",time.Now().Format("2006-01-02 15:04:05.000000") , time.Now().UTC())
		}
	}
}
// 总结&分析
// 上述程序中，代码中 process 函数休眠了 10500 毫秒（10.5 秒），接着把 process successful 写入 ch 信道。
// 在函数 MainCall2 中的 代码行，并发地调用了这个函数。
// 在并发地调用了 process 协程之后，主协程启动了一个无限循环。
// 这个无限循环在每一次迭代开始时，都会先休眠 1000 毫秒（1 秒），
// 然后执行一个 select 操作。在最开始的 10500 毫秒中，由于 process 协程在 10500 毫秒后才会向 ch 信道写入数据，
// 因此 select 语句的第一个 case（即 case v := <-ch:）并未就绪。所以在这期间，程序会执行默认情况，该程序会打印 10 次 no value received。
//
// 在 10.5 秒之后，process 协程会 代码 case v := <-ch 行向 ch 写入 process successful。
// 现在，就可以执行 select 语句的第一个 case 了，程序会打印 received value: process successful，然后程序终止。该程序会输出：