package _021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists_(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := ListNode{}
	cur := &dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1, cur = list1.Next, cur.Next
		} else {
			cur.Next = list2
			list2, cur = list2.Next, cur.Next
		}
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}
