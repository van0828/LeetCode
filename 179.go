package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
}

//给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
//
//示例 1:
//
//输入: [10,2]
//输出: 210
//示例 2:
//
//输入: [3,30,34,5,9]
//输出: 9534330
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) > strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
	})
	res := ""
	for _, num := range nums {
		res = res + strconv.Itoa(num)
	}
	if resNum, _ := strconv.Atoi(res); resNum == 0 {
		return "0"
	}
	return res
}
