package _206

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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur, next *ListNode
	for cur, next = head, head.Next; cur.Next != nil; {
		cur.Next = pre
		pre, cur = cur, next
		next = cur.Next
	}
	cur.Next = pre
	return cur
}
