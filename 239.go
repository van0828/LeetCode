package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{2,4,7}, 2))
}

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口最大值。
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//链接：https://leetcode-cn.com/problems/sliding-window-maximum
func maxSlidingWindow(nums []int, k int) []int {
	var handler func(nums []int) int
	handler = func(nums []int) int {
		max := nums[0]
		for i:=1;i<len(nums);i++{
			if max < nums[i] {
				max = nums[i]
			}
		}
		return max
	}
	result := make([]int, 0)
	if len(nums) == 0 {
		return result
	}
	if k == 1 {
		return nums
	}
	for i := 0; i <= len(nums)-k; i++ {
		result = append(result, handler(nums[i:i+k]))
	}
	return result
}
