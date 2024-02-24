package _549

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

func longestConsecutive(root *TreeNode) (result int) {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		if node.Left == nil && node.Right == nil {
			return 1, 1
		}
		incL, decL := dfs(node.Left)
		incR, decR := dfs(node.Right)

		inc, dec := 1, 1
		incMax, decMax := 1, 1
		if node.Left != nil && node.Val-node.Left.Val == 1 {
			inc = 1 + incL
			incMax += incL
		}
		if node.Right != nil && node.Val-node.Right.Val == -1 {
			dec = 1 + decR
			incMax += decR
		}
		if node.Left != nil && node.Val-node.Left.Val == -1 {
			dec = max(dec, 1+decL)
			decMax += decL
		}
		if node.Right != nil && node.Val-node.Right.Val == 1 {
			inc = max(inc, 1+incR)
			decMax += incR
		}
		result = max(result, incMax, decMax)
		return inc, dec
	}

	result = 1
	dfs(root)
	return
}
