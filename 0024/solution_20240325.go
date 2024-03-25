package _024

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

func swapPairs(head *ListNode) *ListNode {

	reverse := func(start, end *ListNode) *ListNode {
		prev, cur := end, start
		for cur != end {
			next := cur.Next
			cur.Next = prev
			prev, cur = cur, next
		}
		return prev
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		start, end := cur.Next, cur.Next.Next.Next
		cur.Next = reverse(start, end)
		cur = cur.Next.Next
	}

	return dummy.Next
}
