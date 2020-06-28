package main

import "fmt"

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false
//
//链接：https://leetcode-cn.com/problems/valid-palindrome

func main()  {
	fmt.Println(isPalindromeStr("A man, a plan, a canal: Panama"), isPalindromeStr("race a car"), isPalindromeStr("12 34 56, -=654321"))
}

//func isPalindromeStr(s string) bool {
//	for i, j := 0, len(s)-1; i < j; {
//		if !isNumberOrCharacter(s[i]) {
//			i++
//			continue
//		}
//		if !isNumberOrCharacter(s[j]) {
//			j--
//			continue
//		}
//		if !equalsIgnoreCase(s[i], s[j]) {
//			return false
//		} else {
//			i++
//			j--
//		}
//	}
//	return true
//
//}

//func isNumberOrCharacter(x uint8) bool {
//	return (x >= 48 && x <= 57) || isCharacter(x)
//}
//
//func isCharacter(x uint8) bool {
//	return (x >= 65 && x <= 90) || (x >= 97 && x <= 122)
//}
//
//func equalsIgnoreCase(x, y uint8) bool {
//	return x == y || (isCharacter(x) && isCharacter(y) && ((x-y == 32) || (y-x == 32)))
//}
