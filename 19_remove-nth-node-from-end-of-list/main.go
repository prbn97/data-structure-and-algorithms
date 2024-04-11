package main

type headNode struct {
	Val  int
	Next *headNode
}

func main() {
	// Input: head = [1,2,3,4,5], n = 2
	head := &headNode{Val: 1}
	head.Next = &headNode{Val: 2}
	head.Next.Next = &headNode{Val: 3}
	head.Next.Next.Next = &headNode{Val: 4}
	head.Next.Next.Next.Next = &headNode{Val: 5}
	removeNthFromEnd(head, 2)
	// Output: [1,2,3,5]

}

// https://leetcode.com/problems/remove-nth-node-from-end-of-head/description/

func removeNthFromEnd(head *headNode, n int) *headNode {

}
