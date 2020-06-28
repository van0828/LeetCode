package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

//给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
//
//对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
//
//你可以返回任何满足上述条件的数组作为答案。
//sort-array-by-parity-ii
func sortArrayByParityII(A []int) []int {
	j := 1
	for i := 0; i < len(A); i += 2 {
		if A[i]%2 != 0 {
			for ; j < len(A); j += 2 {
				if A[j]%2 != 1 {
					break
				}
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
