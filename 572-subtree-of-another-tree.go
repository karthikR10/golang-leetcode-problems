// https://leetcode.com/problems/subtree-of-another-tree/

// Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

 

// Example 1:


// Input: root = [3,4,5,1,2], subRoot = [4,1,2]
// Output: true
// Example 2:


// Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
// Output: false
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
   if root == nil {
		return false
   }

    sameTree := isSameTree(root, subRoot)
	if sameTree {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}


func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil || s.Val != t.Val {
		return false
	} 

	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}