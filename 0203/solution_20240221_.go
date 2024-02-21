package _203

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements___(head *ListNode, val int) *ListNode {

	var dfs func(node *ListNode, val int) *ListNode
	dfs = func(node *ListNode, val int) *ListNode {
		if node == nil {
			return nil
		}
		if node.Val == val {
			return dfs(node.Next, val)
		}
		node.Next = dfs(node.Next, val)
		return node
	}

	return dfs(head, val)
}
