package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Input: head = [1,2,3,4,5], n = 2
	var head *ListNode

	for i := 5; i >= 1; i-- {
		node := ListNode{
			Val:  i,
			Next: head,
		}

		head = &node
	}

}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// removeNthFromEnd(head, 2)
// Output: [1,2,3,5]

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Criação de um nó cabeça para facilitar a remoção do primeiro elemento
	dummy := &ListNode{Next: head}

	// Inicializar dois ponteiros, leader e follower
	leader := dummy
	follower := dummy

	// Avançar o ponteiro leader por n+1 nós
	for i := 0; i <= n; i++ {
		leader = leader.Next
	}

	// Avançar ambos os ponteiros até que o líder chegue ao final da lista
	for leader != nil {
		leader = leader.Next
		follower = follower.Next
	}

	// Remover o nó referenciado pelo ponteiro follower
	follower.Next = follower.Next.Next

	// Retornar o nó cabeça da lista (dummy.Next)
	return dummy.Next
}
