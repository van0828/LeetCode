package main

func main() {
	
}

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//linked-list-cycle
func hasCycle(head *ListNode) bool {
	var oneStep, twoStep *ListNode
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	oneStep = head.Next
	twoStep = head.Next.Next
	if twoStep == head {
		return true
	}
	for oneStep != twoStep {
		if twoStep.Next == nil || twoStep.Next.Next == nil {
			return false
		} else {
			oneStep = oneStep.Next
			twoStep = twoStep.Next.Next
		}
	}
	return true
}
