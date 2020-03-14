package ch22

// 关闭信道和使用for range 遍历信道
// 数据发送方可以关闭信道,通知接收方这个信道不再又数据发送过来
// 当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭
// 语法 v, ok := <-ch

func Producter(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
