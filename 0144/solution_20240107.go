package _144

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Solution One: Recursively.

func preorderTraversal(root *TreeNode) []int {
	return traverse(root)
}

func traverse(node *TreeNode) (result []int) {
	if node == nil {
		return result
	}
	result = append(result, node.Val)
	result = append(result, traverse(node.Left)...)
	result = append(result, traverse(node.Right)...)
	return
}

// Solution Two: Iteratively.

func preorderTraversal_(root *TreeNode) []int {

	result := make([]int, 0)

	stack := make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		result = append(result, node.Val)
	}

	return result
}
