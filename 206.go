package main

func main() {
	
}

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
// reverse-linked-list
func reverseList(head *ListNode) *ListNode {
	current := head
	var root *ListNode
	for current != nil {
		next := current.Next
		current.Next = root
		root = current
		current = next
	}
	return root
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil { // 如果头为空
//		return head
//	}
//
//	h := getNextNode1(head.Next, head) // n, n-1: 2, 1
//	head.Next = nil
//	return h
//}
//
//// getNextNode node：后节点，preNode：node的前节点
//func getNextNode1(node *ListNode, preNode *ListNode) *ListNode {
//	if node == nil { // 递归终止条件
//		return preNode
//	}
//	m := node.Next
//	node.Next = preNode // 递归条件 n.Next = n-1
//	return getNextNode1(m, node)
//}
