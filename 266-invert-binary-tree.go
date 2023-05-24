// Given the root of a binary tree, invert the tree, and return its root.

// Example 1:

// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
// Example 2:

// Input: root = [2,1,3]
// Output: [2,3,1]

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	temp := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = temp

	return root
}

func main() {

	// Constructing the binary tree
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	inverted := invertTree(root)

	fmt.Println(inverted)
}

// func printTree(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}

// 	fmt.Printf("%d ", root.Val)
// 	printTree(root.Left)
// 	printTree(root.Right)
// }
