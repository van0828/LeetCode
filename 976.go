package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 3, 5, 6, 3, 1}))
}

//给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。
//
//如果不能形成任何面积不为零的三角形，返回 0。
//largest-perimeter-triangle
func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	for i := 2; i < len(A); i++ {
		if A[i-2]+A[i-1] > A[i] && A[i-2]+A[i] > A[i-1] && A[i]+A[i-1] > A[i-2] {
			return A[i-2] + A[i-1] + A[i]
		}
	}
	return 0
}
