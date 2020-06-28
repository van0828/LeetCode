package main

//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//
//示例 1:
//
//输入: 2.00000, 10
//输出: 1024.00000
//示例 2:
//
//输入: 2.10000, 3
//输出: 9.26100
//示例 3:
//
//输入: 2.00000, -2
//输出: 0.25000
//解释: 2-2 = 1/22 = 1/4 = 0.25
//powx-n
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / MyPow(x, -n)
	}
	return MyPow(x, n)
}

func MyPow(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n == 0 {
		return 1
	}
	res := MyPow(x, n/2)
	res *= res
	if n%2 != 0 {
		return res * x
	}


	return res
}
