package _876

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

func middleNode(head *ListNode) *ListNode {

	dummy := &ListNode{Next: head}

	slow, fast := dummy, dummy
	for i := 0; fast != nil; i++ {
		if i%2 == 0 {
			slow = slow.Next
		}
		fast = fast.Next
	}
	return slow
}
