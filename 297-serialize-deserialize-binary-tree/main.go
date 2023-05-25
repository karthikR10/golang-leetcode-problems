// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

// 297. Serialize and Deserialize Binary Tree
// Hard
// 8.8K
// 321
// Companies
// Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

// Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

// Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

// Example 1:

// Input: root = [1,2,3,null,null,4,5]
// Output: [1,2,3,null,null,4,5]
// Example 2:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 104].
// -1000 <= Node.val <= 1000

package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := make([]string, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "N")
			return
		}

		res = append(res, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	i := 0

	var dfs func() *TreeNode

	dfs = func() *TreeNode {
		if vals[i] == "N" {
			i += 1
			return nil
		}

		val, _ := strconv.Atoi(vals[i])
		node := &TreeNode{Val: val}
		i += 1
		node.Left = dfs()
		node.Right = dfs()
		return node
	}

	return dfs()
}

func main() {

}
