package ch24

// 死锁与默认情况

func MainCall3() {
	ch := make(chan string)
	select {
	case <-ch:

	}
}

// 总结&分析
// 运行结果是 fatal error: all goroutines are asleep - deadlock!
// 代码中 创建了一个信道 ch . 在 select 内部，试图读取信道 ch .
// 由于没有 Go 协程向该信道写入数据。
// 因此 select 语句会一直阻塞，导致死锁。该程序会触发运行时 panic,报错信息如下:
// fatal error: all goroutines are asleep - deadlock!
// ch24.MainCall3()
//        /Users/wangyong/workdirs/github/goes/emo/src/ch24/g3.go:8 +0x52


// 如果存在默认情况，就不会发生死锁，因为在没有其他 case 准备就绪时，会执行默认情况。
// 用默认情况重写后 见 go4.go