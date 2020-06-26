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
	if treeA == nil && treeB == nil {
		return true
	}

	if treeA == nil || treeB == nil {
		return false
	}

	if treeA.Val != treeB.Val {
		return false
	}

	return equalTree(treeA.Left, treeB.Left) && equalTree(treeA.Right, treeB.Right)
}

type TreeNodeQueue struct {
	queue []*TreeNode
}

func (t *TreeNodeQueue) Append(n *TreeNode) {
	t.queue = append(t.queue, n)
}

func (t *TreeNodeQueue) PopFront() *TreeNode {
	if t.Len() == 0 {
		return nil
	}
	n := t.queue[0]
	t.queue = t.queue[1:]
	return n
}

func (t *TreeNodeQueue) Len() int {
	return len(t.queue)
}

func isSymmetricII(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queueA := TreeNodeQueue{}
	queueB := TreeNodeQueue{}

	queueA.Append(root.Left)
	queueB.Append(root.Right)

	for ; queueA.Len() != 0 && queueB.Len() != 0; {
		nodeA := queueA.PopFront()
		nodeB := queueB.PopFront()

		if nodeA == nil && nodeB == nil {
			continue
		}

		if nodeA == nil || nodeB == nil {
			return false
		}

		if nodeA.Val != nodeB.Val {
			return false
		}

		queueA.Append(nodeA.Left)
		queueB.Append(nodeB.Right)

		queueA.Append(nodeA.Right)
		queueB.Append(nodeB.Left)

	}

	return queueA.Len() == 0 && queueB.Len() == 0
}
