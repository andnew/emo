package main

import (
	ch "ch25"
	"fmt"
)

func main() {
	fmt.Println("Hello emo!")
	// 第一个示例

	// 第二个示例
	//ch.MainCall2()

	// 第三个示例
	//ch.MainCall3()

	// 第四个示例
	ch.MainCall4()

	// 总结 Mutex vs 信道
	// 通过使用 Mutex 和信道，我们已经解决了竞态条件的问题。那么我们该选择使用哪一个？答案取决于你想要解决的问题。如果你想要解决的问题更适用于 Mutex，那么就用 Mutex。如果需要使用 Mutex，无须犹豫。而如果该问题更适用于信道，那就使用信道。:)
	// 由于信道是 Go 语言很酷的特性，大多数 Go 新手处理每个并发问题时，使用的都是信道。这是不对的。Go 给了你选择 Mutex 和信道的余地，选择其中之一都可以是正确的。
	// 总体说来，当 Go 协程需要与其他协程通信时，可以使用信道。而当只允许一个协程访问临界区时，可以使用 Mutex。
	// 就我们上面解决的问题而言，我更倾向于使用 Mutex，因为该问题并不需要协程间的通信。所以 Mutex 是很自然的选择。
	// 我的建议是去选择针对问题的工具，而别让问题去将就工具。:)
}
