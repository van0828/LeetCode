package main

import "fmt"

//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
//输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//提示：
//
//每个字符串仅由字符 '0' 或 '1' 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 "0" ，就都不含前导零。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-binary
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(addBinary("1010", "1011"), addBinary("11", "1"))
}
func addBinary(a string, b string) string {
	res := ""
	lenA, lenB := len(a)-1, len(b)-1
	var carry uint8
	for lenA >= 0 || lenB >= 0 {
		var va uint8
		if lenA >= 0 {
			va = a[lenA] - '0'
		}
		var vb uint8
		if lenB >= 0 {
			vb = b[lenB] - '0'
		}
		sum := (va + vb + carry) % 2
		res = fmt.Sprintf("%d", sum) + res
		carry = (va + vb + carry) / 2
		lenA--
		lenB--
	}
	if carry != 0 {
		res = "1" + res
	}
	return res
}
