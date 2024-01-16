package _426

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {

	dummy := &Node{}

	var dfs func(node *Node, prev *Node) (last *Node)
	dfs = func(node *Node, prev *Node) (last *Node) {
		if node == nil {
			return nil
		}

		last = dfs(node.Left, prev)
		if last != nil {
			last.Right = node
			node.Left = last
		} else {
			prev.Right = node
			node.Left = prev
		}

		last = dfs(node.Right, node)
		if last == nil {
			last = node
		}
		return
	}

	last := dfs(root, dummy)
	if last == nil {
		return nil
	}

	last.Right = dummy.Right
	dummy.Right.Left = last

	return dummy.Right
}
