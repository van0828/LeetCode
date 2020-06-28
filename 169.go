package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 3, 4}))
}

//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。
//majority-element
//func majorityElement(nums []int) int {
//	result := make(map[int]int, 0)
//	for _, num := range nums {
//		if _, ok := result[num]; !ok {
//			result[num] = 1
//		} else {
//			result[num] = result[num] + 1
//		}
//		if result[num] > len(nums)/2 {
//			return num
//		}
//	}
//	return 0
//}

// 摩尔投票法
func majorityElement(nums []int) int {
	result, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if result == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				result = nums[i]
				count = 1
			}
		}
	}
	return result
}
