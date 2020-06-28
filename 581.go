package main

import "fmt"

func main()  {
	fmt.Println(findUnsortedSubarray([]int{2, 3, 11, 8, 10, 9, 15, 22}))
}

//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//你找到的子数组应是最短的，请输出它的长度。
//
//示例 1:
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//说明 :
//
//输入的数组长度范围在 [1, 10,000]。
//输入的数组可能包含重复元素 ，所以升序的意思是<=。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//输入[2, 3, 11, 8, 10, 9, 15, 22]
//算法流程:
//从左往右遍历数组, 发现{15, 22}是递增的, 同时15比前面所有的元素都大, 说明{15, 22}的位置已经固定了
//从右往左遍历数组, 发现{2, 3}是递增的, 同时3比右边所有的元素都小, 所以{2,3}的位置也已经固定了
//
//经过两遍遍历之后发现{2,3}和{15,22}不用排序了,只需要将{11, 8, 10, 9}这4个元素排序即可, 所以最终输出4
//遍历的过程就是在找哪些元素位置已经固定, 不需要排序了
//
func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := len(nums), -1
	min, max := nums[len(nums)-1], nums[0]

	for index, num := range nums {
		if num >= max {
			max = num
		} else {
			right = index
		}
		fmt.Println(right)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			left = i
		}
	}
	fmt.Println(left)
	if right > left {
		return right - left + 1
	}
	return 0
}
