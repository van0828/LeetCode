package main

import (
	"fmt"
	"sort"
)

func main() {
	result := mergeIntervals([][]int{
		[]int{1, 4},
		[]int{0,0},
		//[]int{8,10},
		//[]int{15,18},
	})
	for _, res := range result {
		fmt.Println(res)
	}
}

//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//merge-intervals
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	result := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else if intervals[i][1] > end {
			end = intervals[i][1]
		}
	}
	result = append(result, []int{start, end})
	return result
}
