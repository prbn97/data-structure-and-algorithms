package main

import "fmt"

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

// list1 = [1,2,4]
// list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Criação de um nó cabeça para a lista resultante
	head := &ListNode{}
	current := head

	// Percorre ambas as listas
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		// Avançar para o próximo nó na lista resultante
		current = current.Next
	}

	// Se uma das listas chegou ao fim, mas a outra não
	// Adiciona o restante dos nós da lista não vazia à lista resultante
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	for head.Next != nil {
		fmt.Println(head.Next.Val)
		head.Next = head.Next.Next

	}
	// Retornar a lista resultante sem o nó cabeça
	return head.Next
}
