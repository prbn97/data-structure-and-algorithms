package main

// https://leetcode.com/problems/merge-two-sorted-lists/
func main() {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	// list1 = [1,2,4]
	// list2 = [1,3,4]
	mergeTwoLists(list1, list2)
	// Output: [1,1,2,3,4,4]

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	currentNodeL1 := list1

	for list2 != nil {
		// if currentNodeL1.Val <= list2.Val {
		// currentNodeL1.Next = list2
		// }

	}

	return list1

}
