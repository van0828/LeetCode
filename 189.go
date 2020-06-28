package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	rotate1(nums1, 4)
	fmt.Println(nums1)

	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate2(nums2, 2)
	fmt.Println(nums2)

	nums3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate3(nums3, 3)
	fmt.Println(nums3)
}

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]
//说明:
//
//尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
//要求使用空间复杂度为 O(1) 的原地算法。
//rotate-array

// 每次向右移动一个 空间O(1) 时间O(kn)
func rotate1(nums []int, k int) {
	if k <= 0 {
		return
	}
	if k >= len(nums) {
		k = k % len(nums)
	}
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

// 翻转 空间O(0) 时间O(n)
func rotate2(nums []int, k int) {
	if k <= 0 {
		return
	}
	if k >= len(nums) {
		k = k % len(nums)
	}
	RotateReverse(nums, 0, len(nums)-1)
	RotateReverse(nums, 0, k-1)
	RotateReverse(nums, k, len(nums)-1)
}

func RotateReverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 每个数字直接移动到最终的位置 空间O(1) 时间O(n)
func rotate3(nums []int, k int) {
	length := len(nums)
	if k <= 0 || k%length == 0 {
		return
	}
	k = k % length
	times := 0
	for i := 0; i < k; i++ {
		if times == length {
			// 每个数字都移动过了 结束
			break
		}
		index := i
		nums[i], nums[(index+k)%length] = nums[(index+k)%length], nums[i]
		index = (index + k) % length
		times++
		for index != i {
			nums[i], nums[(index+k)%length] = nums[(index+k)%length], nums[i]
			index = (index + k) % length
			times++
		}
	}
}
