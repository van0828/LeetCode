package main

import (
	"fmt"
)

//给定一个没有重复数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// permutations
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}

	for i, num := range nums {
		tmp := make([]int, len(nums)-1)
		copy(tmp[:i], nums[:i])
		copy(tmp[i:], nums[i+1:])

		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}
