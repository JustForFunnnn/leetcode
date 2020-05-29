package main

import (
	"container/heap"
	"fmt"
)

/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

	Input:
		[
		  1->4->5,
		  1->3->4,
		  2->6
		]
	Output:
		1->1->2->3->4->4->5->6
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	(*h)[h.Len()-1] = nil
	*h = (*h)[0 : h.Len()-1]
	return x
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	preNode := dummy

	// In case of Nil ListNode
	var notNilLists []*ListNode
	for _, list := range lists {
		if list != nil {
			notNilLists = append(notNilLists, list)
		}
	}
	h := ListNodeHeap(notNilLists)

	heap.Init(&h)
	for ; h.Len() != 0; {
		minimal := heap.Pop(&h)
		minimalListNode := minimal.(*ListNode)
		preNode.Next = minimalListNode

		if minimalListNode.Next != nil {
			heap.Push(&h, minimalListNode.Next)
		}
		preNode = minimalListNode
	}

	return dummy.Next

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

// test case
func main() {
	lists := make([]*ListNode, 0, 3)
	lists = append(lists, newLinkList([]int{1, 4, 5}))
	lists = append(lists, newLinkList([]int{1, 3, 4}))
	lists = append(lists, newLinkList([]int{2, 6}))

	result := mergeKLists(lists)
	fmt.Println("after merge, the link: ")
	for node := result; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Val)
	}
	fmt.Printf("Nil\n")

	fmt.Println(mergeKLists([]*ListNode{nil,}))
}
