package _337

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

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) [2]int
	dfs = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		res := [2]int{}
		res[0] = node.Val + left[1] + right[1]                   // 表示选择当前节点
		res[1] = max(left[0], left[1]) + max(right[0], right[1]) // 表示不选择当前节点
		return res
	}

	res := dfs(root)
	return max(res[0], res[1])
}
