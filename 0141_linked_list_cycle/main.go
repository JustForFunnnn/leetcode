package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slowNode := head
	fastNode := head

	for fastNode != nil{
		if fastNode.Next == nil || fastNode.Next.Next == nil{
			return false
		}

		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next

		if slowNode == fastNode{
			return true
		}
	}


	return false
}
