package _144

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal___(root *TreeNode) []int {

	result := make([]int, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return result
}

func preorderTraversal____(root *TreeNode) (result []int) {

	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		node := nexts[len(nexts)-1]
		nexts = nexts[:len(nexts)-1]
		if node != nil {
			result = append(result, node.Val)
			nexts = append(nexts, node.Right)
			nexts = append(nexts, node.Left)
		}
	}
	return
}
