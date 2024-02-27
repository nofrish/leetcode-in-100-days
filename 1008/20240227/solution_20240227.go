package _0240227

import "math"

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

func bstFromPreorder(preorder []int) *TreeNode {

	idx := 0

	var dfs func(lower, upper int) *TreeNode
	dfs = func(lower, upper int) *TreeNode {
		if idx == len(preorder) {
			return nil
		}
		val := preorder[idx]
		if val < lower || val > upper {
			return nil
		}
		idx++

		node := &TreeNode{val, nil, nil}
		node.Left = dfs(lower, val)
		node.Right = dfs(val, upper)
		return node
	}

	return dfs(math.MinInt, math.MaxInt)
}
