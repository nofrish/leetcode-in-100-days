package _372

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

func longestZigZag(root *TreeNode) (result int) {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		if node.Left == nil && node.Right == nil {
			return 1, 1
		}

		_, ld2 := dfs(node.Left)
		rd1, _ := dfs(node.Right)

		ld, rd := ld2, rd1
		result = max(result, ld, rd)
		return ld + 1, rd + 1
	}

	dfs(root)
	return
}
