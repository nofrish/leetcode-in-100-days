package _112

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) (result bool) {

	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				result = true
			}
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, 0)
	return
}
