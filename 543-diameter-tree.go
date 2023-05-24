// 543. Diameter of Binary Tree
// Easy
// 11.4K
// 708
// Companies
// Given the root of a binary tree, return the length of the diameter of the tree.

// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

// The length of a path between two nodes is represented by the number of edges between them.

// Example 1:

// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
// Example 2:

// Input: root = [1,2]
// Output: 1

package main

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0

	dfs(root, &maxLength)
	return maxLength
}

func dfs(t *TreeNode, maxLength *int) int {
	if t == nil {
		return 0
	}

	left := dfs(t.Left, maxLength)
	right := dfs(t.Right, maxLength)

	*maxLength = max(*maxLength, left+right)

	return max(left, right) + 1

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
