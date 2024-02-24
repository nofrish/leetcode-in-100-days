package _120

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

func maximumAverageSubtree(root *TreeNode) (result float64) {

	var dfs func(node *TreeNode) (sum, count int)
	dfs = func(node *TreeNode) (sum, count int) {
		if node == nil {
			return 0, 0
		}
		if node.Left == nil && node.Right == nil {
			result = max(result, float64(node.Val))
			return node.Val, 1
		}

		lsum, lcount := dfs(node.Left)
		rsum, rcount := dfs(node.Right)

		sum = node.Val + lsum + rsum
		count = 1 + lcount + rcount
		result = max(result, float64(sum)/float64(count))
		return sum, count
	}

	dfs(root)
	return
}
