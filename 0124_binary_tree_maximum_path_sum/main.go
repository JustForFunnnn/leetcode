package main

/*
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections

The path must contain at least one node and does not need to go through the root.

Example 1:
	Input: [1,2,3]

		   1
		  / \
		 2   3

	Output: 6

Example 2:
	Input: [-10,9,20,null,null,15,7]

	   -10
	   / \
	  9  20
		/  \
	   15   7

	Output: 42
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}
	res = root.Val

	getSum(root, &res)
	return res
}

func getSum(root *TreeNode, maxSum *int) int {
	leftChildSum := 0
	if root.Left != nil {
		if sum := getSum(root.Left, maxSum); sum > 0 {
			leftChildSum = sum
		}
	}

	rightChildSum := 0
	if root.Right != nil {
		if sum := getSum(root.Right, maxSum); sum > 0 {
			rightChildSum = sum
		}
	}

	if (rightChildSum + leftChildSum + root.Val) > *maxSum {
		*maxSum = rightChildSum + leftChildSum + root.Val
	}

	if rightChildSum > leftChildSum {
		return rightChildSum + root.Val
	} else {
		return leftChildSum + root.Val
	}
}
