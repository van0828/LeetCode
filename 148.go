package main

func main() {

}

//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//示例 1:
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//sort-list

// 归并排序 1找中点 2递归拆分链表 3合并有序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return divideList(head)
}

func divideList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	l := divideList(head)
	r := divideList(slow)
	return mergeList(l, r)
}

func mergeList(l, r *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			l = l.Next
			cur = cur.Next
		} else {
			cur.Next = r
			r = r.Next
			cur = cur.Next
		}
	}
	if l != nil {
		cur.Next = l
	}
	if r != nil {
		cur.Next = r
	}
	return head.Next
}
