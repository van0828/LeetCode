package main

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//palindrome-linked-list
func isPalindromeNode(head *ListNode) bool {
	num := make([]int, 0)
	for head != nil {
		num = append(num, head.Val)
		head = head.Next
	}
	for i := 0; i < len(num)/2; i++ {
		if num[i] != num[len(num)-1-i] {
			return false
		}
	}
	return true
}
