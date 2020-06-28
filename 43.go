package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("12345678","0"))
}

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//multiply-strings
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	resInts := make([]int, len(num1)+len(num2)+1)
	for i := len(num1) - 1; i >= 0; i-- {
		k := len(num2) + i + 1
		for j := len(num2) - 1; j >= 0; j-- {
			mul := (int((num1[i]-'0')*(num2[j]-'0')) + resInts[k]) % 10
			tmp := (int((num1[i]-'0')*(num2[j]-'0')) + resInts[k]) / 10
			resInts[k] = mul
			k--
			resInts[k] += tmp
		}
	}
	i := 0
	for ; i < len(resInts); i++ {
		if resInts[i] != 0 {
			break
		}
	}
	res := ""
	for ; i < len(resInts); i++ {
		res += strconv.Itoa(resInts[i])
	}
	return res
}
