package main

func main() {

}

//编写一个程序，找到两个单链表相交的起始节点。
//intersection-of-two-linked-lists
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	pA, pB := headA, headB
	for pA != nil {
		lenA++
		pA = pA.Next
	}
	for pB != nil {
		lenB++
		pB = pB.Next
	}
	if lenA >= lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
