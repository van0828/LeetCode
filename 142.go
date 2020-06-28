package main

func main() {

}

// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
// linked-list-cycle-ii
func detectCycle(head *ListNode) *ListNode {
	var oneStep, twoStep *ListNode
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	oneStep = head.Next
	twoStep = head.Next.Next
	if twoStep == head {
		return head
	}
	// 是否有环
	for oneStep != twoStep {
		if twoStep.Next == nil || twoStep.Next.Next == nil {
			return nil
		} else {
			oneStep = oneStep.Next
			twoStep = twoStep.Next.Next
		}
	}
	// 找入环点
	oneStep = head
	for {
		if oneStep == twoStep {
			return oneStep
		} else {
			oneStep = oneStep.Next
			twoStep = twoStep.Next
		}

	}
	return nil
}
