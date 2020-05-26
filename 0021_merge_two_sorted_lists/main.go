package main

import (
	"fmt"
)

/*
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	preNode := dummy

	for ; l1 != nil && l2 != nil; {
		if l1.Val <= l2.Val {
			preNode.Next = l1
			preNode = l1
			l1 = l1.Next
		} else {
			preNode.Next = l2
			preNode = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		preNode.Next = l1
	}

	if l2 != nil {
		preNode.Next = l2
	}

	return dummy.Next
}

// test case
func main() {
	l1 := newLinkList([]int{1, 2, 4})
	l2 := newLinkList([]int{1, 3, 4})

	mergedLink := mergeTwoLists(l1, l2)

	fmt.Println("after merge, the link: ")
	for node := mergedLink; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Val)
	}
	fmt.Printf("Nil\n")
}

func newLinkList(valList []int) (head *ListNode) {
	dummy := &ListNode{Val: -1}
	preNode := dummy

	for _, val := range valList {
		node := &ListNode{Val: val}
		preNode.Next = node
		preNode = node
	}

	return dummy.Next
}
