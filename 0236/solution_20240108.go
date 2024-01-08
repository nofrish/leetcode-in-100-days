package _236

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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	lcs, lancestor := traverse(root.Left, p, q)
	rcs, rancestor := traverse(root.Right, p, q)

	if lcs == 0 {
		return lancestor
	}
	if rcs == 0 {
		return rancestor
	}

	return root
}

func traverse(node, p, q *TreeNode) (state int, ancestor *TreeNode) {
	// state: -1 means nothing is found, 1 means p or q is found, 0 means common ancestor is found
	if node == nil {
		return -1, nil
	}

	lcs, lancestor := traverse(node.Left, p, q)
	rcs, rancestor := traverse(node.Right, p, q)

	// 已经找到LCA，直接返回
	if lcs == 0 {
		return 0, lancestor
	}
	if rcs == 0 {
		return 0, rancestor
	}

	// 本节点就是LCA
	if lcs == 1 && rcs == 1 {
		return 0, node
	}

	// 本节点是p或q，根据情况返回
	if node.Val == p.Val || node.Val == q.Val {
		if lcs == 1 || rcs == 1 {
			return 0, node
		} else {
			return 1, nil
		}
	}

	if lcs == 1 || rcs == 1 {
		return 1, nil
	}
	return -1, nil
}
