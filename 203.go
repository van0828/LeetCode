package main

//删除链表中等于给定值 val 的所有节点。
//
//示例:
//
//输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5
//remove-linked-list-elements
func removeLinkedListElements(head *ListNode, val int) *ListNode {
	root := &ListNode{
		Next: head,
	}
	for tmp := root; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return root.Next
}
