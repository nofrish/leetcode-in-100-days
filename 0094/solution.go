package _094

// solution one: stack + pushLeftNodes
func _inorderTraversal(root *TreeNode) (result []int) {

	stack := make([]*TreeNode, 0, 100)
	pushLeftNodes := func(node *TreeNode) {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}

	pushLeftNodes(root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pushLeftNodes(top.Right)
		result = append(result, top.Val)
	}
	return
}
