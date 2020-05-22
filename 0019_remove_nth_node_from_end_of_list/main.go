package main

import "fmt"

/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:
	Given linked list: 1->2->3->4->5, and n = 2.

	After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
	Given n will always be valid.
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		// n invalid
		return nil
	}

	slowNode, fastNode := head, head

	for i := 0; i < n; i++ {
		if fastNode == nil {
			// n invalid
			return nil
		}
		fastNode = fastNode.Next
	}

	dummy := &ListNode{Val: -1, Next: head}
	preNode := dummy

	for ; fastNode != nil; {
		preNode = slowNode

		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}

	// remove N-th node from end
	preNode.Next = slowNode.Next

	return dummy.Next
}

// test case
func main() {
	linkHead := &ListNode{Val: 1}
	preNode := linkHead
	for i := 2; i <= 5; i++ {
		node := &ListNode{Val: i}
		preNode.Next = node
		preNode = node
	}

	removedLinkHead := removeNthFromEnd(linkHead, 2)

	fmt.Println("after remove, the link: ")
	for node := removedLinkHead; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Val)
	}
	fmt.Printf("end\n")
}
