package ch24

func MainCall7() {
	select {}
}
// 总结&分析
// 除非有 case 执行，select 语句就会一直阻塞着。在这里，select 语句没有任何 case ，因此它会一直阻塞，导致死锁。
// 程序会触发 panic ，
// 执行结果
// fatal error: all goroutines are asleep - deadlock!
// ch24.MainCall7()
///    Users/wangyong/workdirs/github/goes/emo/src/ch24/go7.go:4 +0x20

