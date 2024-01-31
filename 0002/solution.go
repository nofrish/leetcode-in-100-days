package _002

func _addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{0, nil}
	cur := dummy

	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next

		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}

	return dummy.Next
}
