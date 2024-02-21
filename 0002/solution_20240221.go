package _002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers___(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	cur := dummy

	carry := 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		sum := carry + l1.Val + l2.Val
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}
	for ; l1 != nil; l1 = l1.Next {
		sum := carry + l1.Val
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}
	for ; l2 != nil; l2 = l2.Next {
		sum := carry + l2.Val
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
	}

	return dummy.Next
}
