package main

/*
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given
	preorder = [3,9,20,15,7]
	inorder = [9,3,15,20,7]
	Return the following binary tree:

		3
	   / \
	  9  20
		/  \
	   15   7
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// should return error
	if len(preorder) != len(inorder) {
		return nil
	}
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	rootIdxOfInOrder := -1
	for idx, val := range inorder {
		if root.Val == val {
			rootIdxOfInOrder = idx
		}
	}

	root.Left = buildTree(preorder[1:1+rootIdxOfInOrder], inorder[:rootIdxOfInOrder])
	root.Right = buildTree(preorder[1+rootIdxOfInOrder:], inorder[rootIdxOfInOrder+1:])

	return root
}
