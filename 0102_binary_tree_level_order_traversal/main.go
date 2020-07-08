package main

/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:

	Given binary tree [3,9,20,null,null,15,7],
		3
	   / \
	  9  20
		/  \
	   15   7

	return its level order traversal as:
	[
	  [3],
	  [9,20],
	  [15,7]
	]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}

	parentQueue := []*TreeNode{root,}
	for ; len(parentQueue) != 0; {
		childQueue := make([]*TreeNode, 0)

		oneLevelVal := make([]int, 0, len(parentQueue))
		for _, node := range parentQueue {
			oneLevelVal = append(oneLevelVal, node.Val)
			if node.Left != nil {
				childQueue = append(childQueue, node.Left)
			}
			if node.Right != nil {
				childQueue = append(childQueue, node.Right)
			}
		}
		res = append(res, oneLevelVal)

		parentQueue = childQueue
	}

	return res
}
