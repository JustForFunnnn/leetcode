package main

/*
Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

    1
   / \
  2   5
 / \   \
3   4   6
The flattened tree should look like:

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	preNode := &TreeNode{}
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		topIdx := len(stack) - 1
		topNode := stack[topIdx]
		stack = stack[:topIdx]

        if topNode.Right != nil {
			stack = append(stack, topNode.Right)
		}
		if topNode.Left != nil {
			stack = append(stack, topNode.Left)
		}

		preNode.Left = nil
		preNode.Right = topNode
		preNode = topNode
	}
}