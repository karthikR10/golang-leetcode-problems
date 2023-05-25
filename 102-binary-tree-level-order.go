// https://leetcode.com/problems/binary-tree-level-order-traversal/

// 102. Binary Tree Level Order Traversal
// Medium
// 13.1K
// 262
// Companies
// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]
// Example 2:

// Input: root = [1]
// Output: [[1]]
// Example 3:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 2000].
// -1000 <= Node.val <= 1000

package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	q := []*TreeNode{root}

	level := 0

	for len(q) > 0 {
		curLen := len(q)
		result = append(result, []int{})
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]

			result[level] = append(result[level], node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		level++
	}

	return result
}
