package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

//给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器，且 n 的值至少为 2。
//示例:
//
//输入: [1,8,6,2,5,4,8,3,7]
//输出: 49
//container-with-most-water
func maxArea(height []int) int {
	res, l, r := 0, 0, len(height)-1
	for l < r {
		res = int(math.Max(float64(res), math.Min(float64(height[l]), float64(height[r]))*float64(r-l)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
