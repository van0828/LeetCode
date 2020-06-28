package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{2,1,2,5,3,2}))
}

// 在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。
// 返回重复了 N 次的那个元素。

// n-repeated-element-in-size-2n-array
func repeatedNTimes(A []int) int {
	var result int
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] ||
			(i < len(A)-2 && A[i] == A[i+2]) ||
			(i < len(A)-3 && A[i] == A[i+3]) {
			result = A[i]
			break
		}
	}
	return result
}
