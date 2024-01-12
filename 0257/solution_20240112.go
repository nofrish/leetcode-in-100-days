package _257

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	return traverse(root, []string{})
}

func traverse(node *TreeNode, prevs []string) []string {
	if node.Left == nil && node.Right == nil {
		if len(prevs) > 0 {
			return []string{fmt.Sprintf("%s->%d", strings.Join(prevs, "->"), node.Val)}
		} else {
			return []string{fmt.Sprintf("%d", node.Val)}
		}
	}
	result := make([]string, 0)
	if node.Left != nil {
		result = append(result, traverse(node.Left, append(prevs, strconv.Itoa(node.Val)))...)
	}
	if node.Right != nil {
		result = append(result, traverse(node.Right, append(prevs, strconv.Itoa(node.Val)))...)
	}
	return result
}
