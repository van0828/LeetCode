package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"ower", "flow", "ight"}))
}

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。
// longest-common-prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	result := ""
	for i := 0; i < len(strs[0]); i++ {
		allHas := true
		// 判断每个位置字符是否都相等 注意不要越界
		for _, str := range strs {
			if i > len(str)-1 || str[i] != strs[0][i] {
				allHas = false
				break
			}
		}
		if allHas {
			result = result + string(strs[0][i])
		} else {
			break
		}
	}
	return result
}
