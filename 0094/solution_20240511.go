package _094

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal___(root *TreeNode) (result []int) {

	nexts := make([]*TreeNode, 0)
	pushLeft := func(node *TreeNode) {
		if node != nil {
			nexts = append(nexts, node)
			node = node.Left
		}
	}

	pushLeft(root)
	for len(nexts) != 0 {
		top := nexts[len(nexts)-1]
		nexts = nexts[:len(nexts)-1]
		pushLeft(top.Right)
		result = append(result, top.Val)
	}
	return
}
