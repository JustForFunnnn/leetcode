package main

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

		1
	   / \
	  2   2
	 / \ / \
	3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

		1
	   / \
	  2   2
	   \   \
	   3    3

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	mirrorRoot := mirrorTree(root)
	return equalTree(mirrorRoot, root)
}

func mirrorTree(root *TreeNode) (newTree *TreeNode) {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{
		Val:   root.Val,
		Left:  mirrorTree(root.Right),
		Right: mirrorTree(root.Left),
	}

	return newRoot
}

func equalTree(treeA, treeB *TreeNode) bool {
	if treeA == nil && treeB == nil{
		return true
	}

	if treeA == nil || treeB == nil{
		return false
	}

	if treeA.Val != treeB.Val {
		return false
	}

	return equalTree(treeA.Left, treeB.Left) && equalTree(treeA.Right, treeB.Right)
}
