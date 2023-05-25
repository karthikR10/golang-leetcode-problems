// https://leetcode.com/problems/binary-tree-maximum-path-sum/

// 124. Binary Tree Maximum Path Sum
// Hard
// 14.2K
// 650
// Companies
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

// The path sum of a path is the sum of the node's values in the path.

// Given the root of a binary tree, return the maximum path sum of any non-empty path.

// Example 1:

// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
// Example 2:

// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

// Constraints:

// The number of nodes in the tree is in the range [1, 3 * 104].
// -1000 <= Node.val <= 1000
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var res = []int{math.MinInt32}
	dfs(root, res)
	return res[0]
}

func dfs(node *TreeNode, arr []int) int {
	if node == nil {
		return 0
	}

	left := max(0, dfs(node.Left, arr))
	right := max(0, dfs(node.Right, arr))

	arr[0] = max(arr[0], node.Val+left+right)

	return node.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(maxPathSum(root))
}
