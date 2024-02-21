package _002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers____(l1 *ListNode, l2 *ListNode) *ListNode {

	var dfs func(l1, l2 *ListNode, carry int) *ListNode
	dfs = func(l1, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil {
			if carry == 0 {
				return nil
			} else {
				return &ListNode{carry, nil}
			}
		}

		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		return &ListNode{
			Val:  sum % 10,
			Next: dfs(l1, l2, sum/10),
		}
	}

	return dfs(l1, l2, 0)
}
