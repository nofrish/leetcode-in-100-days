package _094

func inorderTraversal(root *TreeNode) []int {
	return traverse(root)
}

func traverse(node *TreeNode) []int {
	result := make([]int, 0)
	if node == nil {
		return result
	}
	result = append(result, traverse(node.Left)...)
	result = append(result, node.Val)
	result = append(result, traverse(node.Right)...)
	return result
}
