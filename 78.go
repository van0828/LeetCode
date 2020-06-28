package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7, 1, 8}))
}

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//subsets
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})
	for i := 0; i < len(nums); i++ {
		lenRes := len(result)
		for j := 0; j < lenRes; j++ {
			//tmp := result[j]
			tmp := make([]int, len(result[j]))
			copy(tmp, result[j])
			tmp = append(tmp, nums[i])
			result = append(result, tmp)
		}
	}
	return result
}
