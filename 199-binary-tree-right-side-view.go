// https://leetcode.com/problems/binary-tree-right-side-view/

// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

// Example 1:

// Input: root = [1,2,3,null,5,null,4]
// Output: [1,3,4]
// Example 2:

// Input: root = [1,null,3]
// Output: [1,3]
// Example 3:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
package main

func rightSideView(root *TreeNode) []int {
	var result []int
	q := []*TreeNode{root}
	if root == nil {
		return result
	}

	for len(q) > 0 {
		var rightTree *TreeNode
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			node := q[0]
			// pop from the queue
			q = q[1:]

			if node != nil {
				rightTree = node
				q = append(q, rightTree.Left)
				q = append(q, rightTree.Right)
			}

		}

		if rightTree != nil {
			result = append(result, rightTree.Val)
		}
	}
	return result
}
