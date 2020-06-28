package main

import "fmt"

func main() {
	nums1 := []int{1,2,3,0,0,0}
	merge(nums1,3,[]int{2,5,6},3)
	fmt.Println(nums1)
}

//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//示例:
//
//输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]
// merge-sorted-array
func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	for i, j := m-1, n-1; i >= 0 || j >= 0; {
		if i < 0 {
			nums1[pos] = nums2[j]
			pos--
			j--
			continue
		}
		if j < 0 {
			nums1[pos] = nums1[i]
			pos--
			i--
			continue
		}
		if nums1[i] >= nums2[j] {
			nums1[pos] = nums1[i]
			pos--
			i--
		} else {
			nums1[pos] = nums2[j]
			pos--
			j--
		}
	}
}
