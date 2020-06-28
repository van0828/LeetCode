package main

//给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
//
//说明：不要使用任何内置的库函数，如  sqrt。
//valid-perfect-square
func isPerfectSquare(num int) bool {
	sumnum := 1
	for num > 0 {
		num -= sumnum
		sumnum += 2
	}
	return num == 0
}
