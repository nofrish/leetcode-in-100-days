package _002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers__(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := ListNode{}
	cur := &dummy

	carry := 0
	for ; l1 != nil && l2 != nil; l1, l2, cur = l1.Next, l2.Next, cur.Next {
		sum := l1.Val + l2.Val + carry
		cur.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		carry = sum / 10
	}

	left := l1
	if l2 != nil {
		left = l2
	}
	for ; left != nil; left, cur = left.Next, cur.Next {
		sum := left.Val + carry
		cur.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		carry = sum / 10
	}

	if carry > 0 {
		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return dummy.Next
}
