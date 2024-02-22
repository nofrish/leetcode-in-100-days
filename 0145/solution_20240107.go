package _145

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

func postorderTraversal(root *TreeNode) (result []int) {

	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		node := nexts[len(nexts)-1]
		nexts = nexts[:len(nexts)-1]
		if node != nil {
			result = append(result, node.Val)
			nexts = append(nexts, node.Left)
			nexts = append(nexts, node.Right)
		}
	}

	reverse(result)
	return
}
