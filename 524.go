package main

import "fmt"

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}

//给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。
//
//示例 1:
//
//输入:
//s = "abpcplea", d = ["ale","apple","monkey","plea"]
//
//输出:
//"apple"
//示例 2:
//
//输入:
//s = "abpcplea", d = ["a","b","c"]
//
//输出:
//"a"
//longest-word-in-dictionary-through-deleting
func isMatch(source, target string) bool {
	for si, ti := 0, 0; si < len(source); si++ {
		if len(target[ti:]) > len(source[si:]) {
			return false
		}
		if target[ti] == source[si] {
			if ti == len(target)-1 {
				return true
			}
			ti++
			continue
		}
	}
	return false
}

func findLongestWord(s string, d []string) string {
	answer := ""
	for _, dw := range d {
		if !isMatch(s, dw) {
			continue
		}
		if len(dw) > len(answer) {
			answer = dw
		}
		if len(dw) == len(answer) && dw < answer {
			answer = dw
		}
	}
	return answer
}
