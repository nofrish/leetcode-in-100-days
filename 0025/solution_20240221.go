package _025

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

func reverseKGroup(head *ListNode, k int) *ListNode {

	reverse := func(node *ListNode, k int) *ListNode {
		var pre, cur *ListNode = nil, node
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre, cur = cur, next
		}
		return pre
	}

	dummy := &ListNode{0, head}
	cur, next := dummy, dummy.Next

	for {
		for j := 0; j < k; j++ {
			if next == nil {
				return dummy.Next
			}
			next = next.Next
		}
		cur.Next = reverse(cur.Next, k)
		for j := 0; j < k; j++ {
			cur = cur.Next
		}
		cur.Next = next
	}

	return dummy.Next
}
