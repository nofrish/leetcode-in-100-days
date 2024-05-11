package _094

// solution one: stack + pushLeft
func _inorderTraversal(root *TreeNode) (result []int) {

	stack := make([]*TreeNode, 0, 100)
	pushLeft := func(node *TreeNode) {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}

	pushLeft(root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pushLeft(top.Right)
		result = append(result, top.Val)
	}
	return
}
