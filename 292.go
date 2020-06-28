package main

import "fmt"

func main() {
	fmt.Println(canWinNim(111))
}

//你和你的朋友，两个人一起玩 Nim游戏：桌子上有一堆石头，每次你们轮流拿掉 1 - 3 块石头。 拿掉最后一块石头的人就是获胜者。你作为先手。
//
//你们是聪明人，每一步都是最优解。 编写一个函数，来判断你是否可以在给定石头数量的情况下赢得游戏。
//nim-game
func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}
	return true
}
