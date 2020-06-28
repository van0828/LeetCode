package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,4,1}))
}

//给定一个整数数组，判断是否存在重复元素。
//
//如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
//
//示例 1:
//
//输入: [1,2,3,1]
//输出: true
// contains-duplicate
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i:=0;i<len(nums)-1;i++{
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
