// 110. Balanced Binary Tree
// Easy
// 8.9K
// 503
// Companies
// Given a binary tree, determine if it is
// height-balanced
// .

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: true
// Example 2:

// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false
// Example 3:

// Input: root = []
// Output: true

// Constraints:

// The number of nodes in the tree is in the range [0, 5000].
// -104 <= Node.val <= 104

package main

import "math"

type balancedTree struct {
	Balance bool
	Height  int
}

func isBalanced(root *TreeNode) bool {
	return dfsTree(root).Balance
}

func dfsTree(t *TreeNode) balancedTree {
	if t == nil {
		return balancedTree{true, 0}
	}

	left := dfsTree(t.Left)
	right := dfsTree(t.Right)

	balheight := int(math.Abs(float64(left.Height)-float64(right.Height))) <= 1
	balanced := (left.Balance && right.Balance && balheight)

	return balancedTree{balanced, 1 + max(left.Height, right.Height)}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
