package _144

var result = make([]int, 0)

func preorderTraversal__(root *TreeNode) []int {
	result = result[0:0:0]
	dfs(root)
	return result
}

func dfs(node *TreeNode) {
	if node != nil {
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
}
