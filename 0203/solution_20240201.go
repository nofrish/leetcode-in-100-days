package _203

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements_(head *ListNode, val int) *ListNode {

	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for ; fast != nil; fast = fast.Next {
		if fast.Val != val {
			slow = slow.Next
			slow.Val = fast.Val
		}
	}
	slow.Next = nil
	return dummy.Next
}
