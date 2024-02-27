package _297

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	sb := strings.Builder{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("#,")
			return
		}
		sb.WriteString(fmt.Sprintf("%d,", node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data[:len(data)-1], ",")
	idx := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if idx == len(parts)-1 {
			return nil
		}

		part := parts[idx]
		idx++

		if part == "#" {
			return nil
		}

		val, _ := strconv.Atoi(part)
		node := &TreeNode{val, nil, nil}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
