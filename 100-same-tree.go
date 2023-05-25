// 100. Same Tree
// Easy
// 9.4K
// 186
// Companies
// Given the roots of two binary trees p and q, write a function to check if they are the same or not.

// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

 

// Example 1:


// Input: p = [1,2,3], q = [1,2,3]
// Output: true
// Example 2:


// Input: p = [1,2], q = [1,null,2]
// Output: false
package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q == nil) {
		return true
	}

	if (p == nil || q == nil || p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}