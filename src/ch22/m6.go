package ch22

//该程序会计算一个数中每一位的平方和与立方和，然后把平方和与立方和相加并打印出来。
// 我们将把这段代码抽离出来，放在一个单独的函数里，然后并发地调用它。
//
func digists(number int, dchnl chan int) {

	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

// 计算一个数种每一位的平方
func CalcSquares2(number int, squareop chan int) {

	sum := 0
	dch := make(chan int)
	go digists(number, dch)

	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

// 计算每一个数中每位数字的立方
func CalcCubes2(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digists(number, dch)

	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

// 总结
// 抽离 函数 CalcSquares2 和 CalcCubes2 的 公用代码，采用 信道的 关闭 机制
//