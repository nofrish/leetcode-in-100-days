package _092

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

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	// 我的经典实现
	reverse := func(start, end *ListNode) *ListNode {
		pre, cur := end, start
		for cur != end {
			next := cur.Next
			cur.Next = pre
			pre, cur = cur, next
		}
		return pre
	}

	dummy := &ListNode{0, head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	start, end := pre.Next, pre.Next
	for i := 0; i <= right-left; i++ {
		end = end.Next
	}

	pre.Next = reverse(start, end)
	return dummy.Next
}
