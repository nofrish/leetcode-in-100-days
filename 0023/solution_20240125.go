package _023

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

func mergeKLists(lists []*ListNode) *ListNode {

	dummy := ListNode{0, nil}
	cur := &dummy

	done := false
	for !done {
		max := -1
		done = true
		for i := range lists {
			if lists[i] == nil {
				continue
			}
			done = false
			if max == -1 || lists[max].Val > lists[i].Val {
				max = i
			}
		}

		if !done {
			cur.Next = lists[max]
			cur = cur.Next
			lists[max] = lists[max].Next
		}
	}

	return dummy.Next
}
