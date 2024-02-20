package _669

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

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	dummy := ListNode{0, list1}
	cur, left, right := &dummy, &dummy, &dummy
	for i := 0; cur != nil; cur, i = cur.Next, i+1 {
		if i == a {
			left = cur
		}
		if i == b {
			right = cur.Next.Next
			break
		}
	}

	head, tail := list2, list2
	for tail.Next != nil {
		tail = tail.Next
	}

	left.Next = head
	tail.Next = right

	return dummy.Next
}
