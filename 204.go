package main

import "fmt"

//统计所有小于非负整数 n 的质数的数量。
//
//示例:
//
//输入: 10
//输出: 4
//解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 11 13 17  。
func main() {
	fmt.Println(countPrimes(10))
}
func countPrimes(n int) int {
	signs := make([]bool, n)
	for index := range signs {
		signs[index] = true
	}
	for i := 2; i < n/i+1; i++ {
		if signs[i] {
			for j := 2 * i; j < n; j += i {
				signs[j] = false
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if signs[i] {
			count++
		}
	}
	fmt.Println(signs)
	return count
}
