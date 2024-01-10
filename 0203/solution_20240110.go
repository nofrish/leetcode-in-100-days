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

// Solution 2. Iteratively
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

// Solution 2. Recursively
func removeElementsRecursively(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	return recursivelyRemove(dummy, dummy.Next, val)
}

func recursivelyRemove(prev, cur *ListNode, val int) *ListNode {
	if cur == nil {
		return nil
	}
	if cur.Val != val {
		cur.Next = recursivelyRemove(cur, cur.Next, val)
		return cur
	} else {
		return recursivelyRemove(prev, cur.Next, val)
	}
}
