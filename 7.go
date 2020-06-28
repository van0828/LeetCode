package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-120))
}

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
//输入: -123
//输出: -321
//示例 3:
//
//输入: 120
//输出: 21

//reverse-integer
func reverse(x int) int {
	nums := make([]int, 0)
	for {
		if x < 10 && x > -10 {
			nums = append(nums, x)
			break
		}
		nums = append(nums, x%10)
		x = x / 10

	}
	fmt.Println(nums)
	result := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		result = result + nums[i]*int(math.Pow10(len(nums)-i-1))
		if result > (1<<31)-1 || result < -1<<31 {
			return 0
		}
	}
	return result
}
