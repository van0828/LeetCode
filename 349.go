package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4]
//intersection-of-two-arrays
func intersection(nums1 []int, nums2 []int) []int {
	resMap := make(map[int]int, 0)
	resSlice := make([]int, 0)
	for _, num := range nums1 {
		resMap[num] = 0
	}
	for _, num := range nums2 {
		val, ok := resMap[num]
		if ok && val == 0 {
			resMap[num]++
			resSlice = append(resSlice, num)
		}
	}
	return resSlice
}
