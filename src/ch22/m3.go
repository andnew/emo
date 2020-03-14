package ch22

//该程序会计算一个数中每一位的平方和与立方和，然后把平方和与立方和相加并打印出来。

// 计算一个数种每一位的平方
func CalcSquares(number int, squareop chan int) {

	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

// 计算每一个数中每位数字的立方
func CalcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}
