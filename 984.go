package main

import (
	"fmt"
)

func main() {
	fmt.Println(strWithout3a3b(4, 1))
}

// 给定两个整数 A 和 B，返回任意字符串 S，要求满足：
//
//S 的长度为 A + B，且正好包含 A 个 'a' 字母与 B 个 'b' 字母；
//子串 'aaa' 没有出现在 S 中；
//子串 'bbb' 没有出现在 S 中。

// string-without-aaa-or-bbb
func strWithout3a3b(A int, B int) string {
	if (A >= B && B*2+2 < A) || (A < B && A*2+2 < B) {
		// 不合理输入
		return "非法"
	}
	out := ""
	a := "a"
	b := "b"

	for A > 0 || B > 0 {
		appendA := false
		if len(out) >= 2 && out[len(out)-1] == out[len(out)-2] {
			if out[len(out)-1] == 'b' {
				appendA = true
			}
		} else if A >= B {
			appendA = true
		}
		if appendA {
			out += a
			A--
		} else {
			out += b
			B--
		}
	}
	return out
}
