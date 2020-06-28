package main

import "fmt"

/*
Given a binary tree, return the in-order traversal of its nodes' values.

Example:

	Input: [1,null,2,3]
	   1
		\
		 2
		/
	   3

	Output: [1,3,2]

Follow up: Recursive solution is trivial, could you do it iteratively?
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(root *TreeNode) (inOrderValList []int) {
	if root == nil {
		return nil
	}

	inOrderValList = append(inOrderValList, inOrderTraversal(root.Left)...)
	inOrderValList = append(inOrderValList, root.Val)
	inOrderValList = append(inOrderValList, inOrderTraversal(root.Right)...)

	return inOrderValList
}

func inOrderTraversalII(root *TreeNode) (inOrderValList []int) {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	node := root

	for ; node != nil || len(stack) != 0; {
		for ; node != nil; {
			stack = append(stack, node)
			node = node.Left
		}

		inNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		inOrderValList = append(inOrderValList, inNode.Val)

		if inNode.Right != nil {
			node = inNode.Right
		}
	}

	return inOrderValList
}
