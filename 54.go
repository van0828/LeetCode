package main

import "fmt"

func main() {
	var arr_6 = [][]int{{1, 2, 3}, {4,5, 6}, {7,8,9}}
	fmt.Println(spiralOrder(arr_6))
}

//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
//示例 1:
//
//输入:
//[
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//输出: [1,2,3,6,9,8,7,4,5]
//示例 2:
//
//输入:
//[
//  [1, 2, 3, 4],
//  [5, 6, 7, 8],
//  [9,10,11,12]
//]
//输出: [1,2,3,4,8,12,11,10,9,5,6,7]
//spiral-matrix

func spiralOrder(matrix [][]int) []int {
	// 一层四条边 依次往下走 注意最后仅剩单列和单行的情况
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	res := make([]int, 0)
	x1, y1, x2, y2 := 0, 0, m-1, n-1
	for x1 <= x2 && y1 <= y2 {
		for y := y1; y <= y2; y++ {
			res = append(res, matrix[x1][y])
		}
		for x := x1 + 1; x <= x2; x++ {
			res = append(res, matrix[x][y2])
		}
		if x1 < x2 && y1 < y2 {
			for y := y2 - 1; y > y1; y-- {
				res = append(res, matrix[x2][y])
			}
			for x := x2; x > x1; x-- {
				res = append(res, matrix[x][y1])
			}
		}
		x1++
		y1++
		x2--
		y2--
	}
	return res
}
