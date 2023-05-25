package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var val int
	inOrderTraverse(root, &count, &val, k)
	return val
}

func inOrderTraverse(root *TreeNode, count *int, val *int, k int) {
	if root == nil {
		return
	}

	inOrderTraverse(root.Left, count, val, k)
	*count++
	if *count == k {
		*val = root.Val
		return
	}
	inOrderTraverse(root.Right, count, val, k)
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	k := 2
	fmt.Println(kthSmallest(root, k))
}

//    3
// 1     4
//   2

// [1, 2, 3, 4]
