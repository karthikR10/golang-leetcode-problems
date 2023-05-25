// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// 230. Kth Smallest Element in a BST
// Medium
// 9.7K
// 175
// Companies
// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

// Example 1:

// Input: root = [3,1,4,null,2], k = 1
// Output: 1
// Example 2:

// Input: root = [5,3,6,2,4,null,null,1], k = 3
// Output: 3

// Constraints:

// The number of nodes in the tree is n.
// 1 <= k <= n <= 104
// 0 <= Node.val <= 104
package main

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var val int
	inOrder(root, &count, &val, k)
	return val
}

func inOrder(root *TreeNode, count, val *int, k int) {
	if root == nil {
		return
	}

	inOrder(root.Left, count, val, k)
	*count++
	if *count == k {
		*val = root.Val
		return
	}
	inOrder(root.Right, count, val, k)
}

//    3
// 1     4
//   2

// [1, 2, 3, 4] 

// k = 2
// output = 2