package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	sl1, sl2 := []int{2, 4, 3}, []int{5, 6, 4}
	l1, l2 := &ListNode{}, &ListNode{}

	tmp1, tmp2 := l1, l2
	for i, v := range sl1 {
		tmp1.Val = v
		if i < len(sl1)-1 {
			tmp1.Next = &ListNode{}
			tmp1 = tmp1.Next
		}
	}

	for i, v := range sl2 {
		tmp2.Val = v
		if i < len(sl2)-1 {
			tmp2.Next = &ListNode{}
			tmp2 = tmp2.Next
		}
	}
	r := addTwoNumbers(l1,l2)
	fmt.Println(r)
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

//add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1d := l1
	l2d := l2
	carry := 0
	preHead := &ListNode{
		Val:0,
		Next:nil,
	}
	curr := preHead
	for l1d != nil || l2d != nil {
		var x, y int
		if l1d != nil {
			x = l1d.Val
		}
		if l2d != nil {
			y = l2d.Val
		}
		sum := carry + x + y
		carry = sum/10
		node := &ListNode{
			Val:sum%10,
		}
		curr.Next = node
		curr = curr.Next
		if l1d != nil {
			l1d = l1d.Next
		}
		if l2d != nil {
			l2d = l2d.Next
		}
	}
	if carry > 0 {
		node := ListNode{
			Val:carry,
		}
		curr.Next = &node
	}
	return preHead.Next
}
