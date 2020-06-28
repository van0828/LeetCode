package main

import "fmt"

func main() {
	fmt.Println(generatePascalsTriangle(5))
}

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]
//
//pascals-triangle

func generatePascalsTriangle(numRows int) [][]int {
	// 没啥说的 肩部两数之和
	if numRows == 0 {
		return nil
	}
	result := make([][]int, 1)
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for k := 1; k < i; k++ {
			tmp[k] = result[i-1][k-1] + result[i-1][k]
		}
		result = append(result, tmp)
	}
	return result
}
