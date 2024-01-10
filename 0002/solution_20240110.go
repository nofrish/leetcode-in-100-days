package _002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	r := &ListNode{0, nil}

	cur := r
	carry := 0

	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		sum := carry + l1.Val + l2.Val
		if sum >= 10 {
			carry = 1
			cur.Next = &ListNode{sum - 10, nil}
		} else {
			carry = 0
			cur.Next = &ListNode{sum, nil}
		}
		cur = cur.Next
	}

	for ; l1 != nil; l1 = l1.Next {
		sum := carry + l1.Val
		if sum >= 10 {
			carry = 1
			cur.Next = &ListNode{sum - 10, nil}
		} else {
			carry = 0
			cur.Next = &ListNode{sum, nil}
		}
		cur = cur.Next
	}

	for ; l2 != nil; l2 = l2.Next {
		sum := carry + l2.Val
		if sum >= 10 {
			carry = 1
			cur.Next = &ListNode{sum - 10, nil}
		} else {
			carry = 0
			cur.Next = &ListNode{sum, nil}
		}
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{1, nil}
	}

	return r.Next
}

func addTwoNumbers_(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{0, nil}
	tail := head

	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		tail.Next = &ListNode{sum % 10, nil}
		tail = tail.Next

		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}

	return head.Next
}
