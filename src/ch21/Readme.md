package main

import (
	"ch21"
	"config"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello emo!")
	//config.LoadConfig()
	config.LoadConfig()

	// 第一个示例
	go ch21.Hello()
	//总结
	//该程序只会输出文本 main function。我们启动的 Go 协程究竟出现了什么问题？要理解这一切，我们需要理解两个 Go 协程的主要性质。
	//
	//启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。
	//在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
	//如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。

	//示例2 等待协程的结果
	time.Sleep(1 * time.Second)

	//示例3 多个协程
	go ch21.Number()
	go ch21.Alphabets()
	time.Sleep(3000 * time.Millisecond)

	log.Println("主协程======")

}
