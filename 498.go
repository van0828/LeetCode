package main

import "fmt"

func main() {
	//var arr_6 = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	//var arr_7 = [][]int{{2}, {3}}
	var arr_8 = [][]int{{2, 5}, {8, 4}, {0, -1}}
	fmt.Println(findDiagonalOrder(arr_8))
}

//给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
//
//
//
//示例:
//
//输入:
//[
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//
//输出:  [1,2,4,7,5,3,6,8,9]
//diagonal-traverse
func findDiagonalOrder(matrix [][]int) []int {
	// 对角线遍历 一条一条看 部分需要反向 注意边界情况
	m := len(matrix)
	if m == 0 {
		return nil
	} else if m == 1 {
		return matrix[0]
	}
	n := len(matrix[0])
	out := make([]int, m*n)
	if n == 1 {
		for i := 0; i < m; i++ {
			out[i] = matrix[i][0]
		}
		return out
	}
	fmt.Printf("m:%d n:%d\n", m, n)
	outPos := 0
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 {
			x, y := 0, i
			if y >= n {
				x, y = i-n+1, n-1
			}
			for {
				fmt.Printf("a x:%d y:%d\n", x, y)
				out[outPos] = matrix[x][y]
				outPos++
				if y == 0 || x == m-1 {
					break
				}
				x++
				y--
			}
		} else {
			x, y := i, 0
			if x >= m {
				x, y = m-1, i-m+1
			}
			for {
				fmt.Printf("b x:%d y:%d\n", x, y)
				out[outPos] = matrix[x][y]
				outPos++
				if x == 0 || y == n-1 {
					break
				}
				x--
				y++
			}
		}
	}
	return out
}
