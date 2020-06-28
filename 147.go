package main

//对链表进行插入排序。
//插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
//每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。
//插入排序算法：
//插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//重复直到所有输入数据插入完为止。
//示例 1：
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2：
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5
//insertion-sort-list
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head

	for head != nil  && head.Next != nil {
		if head.Next.Val >= head.Val {
			//不需要动
			// fmt.Println("不需要动")
			head = head.Next
			continue
		} else {
			//for 循环 从头往后找，找到第一个比当前node大的那个node 插入到这个node前面
			pre := dummy
			for pre.Next != nil {
				if head.Next.Val < pre.Next.Val {
					//证明找到啦 head.Next这个node要换到 pre 这个node之后 pre.Next 这个node 之前
					cur := head.Next
					head.Next = cur.Next
					cur.Next = pre.Next
					pre.Next = cur
					break
				}
				pre = pre.Next
			}
		}
	}

	return dummy.Next
}
