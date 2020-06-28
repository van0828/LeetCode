package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 3sum
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == -nums[i] {
				result = append(result, []int{nums[left], nums[right], nums[i]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] < -nums[i] {
				left++
			} else if nums[left]+nums[right] > -nums[i] {
				right--
			}
		}

	}
	return result
}
