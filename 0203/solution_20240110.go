package _203

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

func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{0, head}
	for curser := dummy; curser.Next != nil; {
		if curser.Next.Val == val {
			curser.Next = curser.Next.Next
		} else {
			curser = curser.Next
		}
	}
	return dummy.Next
}
