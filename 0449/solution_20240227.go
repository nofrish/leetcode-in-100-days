package _449

import (
	"fmt"
	"math"
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	parts := strings.Split(data[:len(data)], ",")
	idx := 0

	var dfs func(lower, upper int) *TreeNode
	dfs = func(lower, upper int) *TreeNode {
		if idx == len(parts) || parts[idx] == "" {
			return nil
		}

		val, _ := strconv.Atoi(parts[idx])
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

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
