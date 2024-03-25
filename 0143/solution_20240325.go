package _143

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

func reorderList(head *ListNode) {

	reverse := func(start, end *ListNode) *ListNode {
		pre, cur := end, start
		for cur != end {
			next := cur.Next
			cur.Next = pre
			pre, cur = cur, next
		}
		return pre
	}

	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}

	half := count / 2
	cur := head
	for i := 0; i < half; i++ {
		cur = cur.Next
	}

	list1 := head
	list2 := reverse(cur.Next, nil)
	cur.Next = nil

	dummy := &ListNode{}
	cur = dummy

	for list1 != nil || list2 != nil {
		if list1 != nil {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		}
		if list2 != nil {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}

	head = dummy.Next
}
