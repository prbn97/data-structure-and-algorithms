// https://leetcode.com/problems/add-two-numbers/
package main

import "fmt"

func (l ListNode) printListData() {
	for l.Next != nil { // Percorre a lista até o último nó
		fmt.Printf("%d ", l.Val)
		l = *l.Next
	}
	fmt.Printf("%d\n", l.Val) // Imprime o valor do último nó
}

func main() {
	// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	// Output: [8,9,9,9,0,0,0,1]
	l1 := &ListNode{Val: 9}
	l1.addValueToList(9)
	l1.addValueToList(9)
	l1.addValueToList(9)
	l1.addValueToList(9)
	l1.addValueToList(9)
	l1.addValueToList(9)
	l1.printListData()
	l2 := &ListNode{Val: 9}
	l2.addValueToList(9)
	l2.addValueToList(9)
	l2.addValueToList(9)
	l2.printListData()
	sumList := addTwoNumbers(l1, l2)
	sumList.printListData()

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) addValueToList(value int) {
	for l.Next != nil { // Percorre a lista até o último nó
		l = l.Next
	}
	newNode := &ListNode{Val: value} // Cria um novo nó com o valor fornecido
	l.Next = newNode                 // Define o próximo do último nó como o novo nó criado
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Nó dummy para simplificar o código
	current := dummy     // Nó atual para construir a nova lista
	carry := 0           // Variável de transporte para armazenar o dígito de transporte

	// Percorre ambas as listas até que ambas sejam nulas e não haja mais transporte
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // Inicializa a soma com o dígito de transporte

		if l1 != nil { //Se o nó atual de l1 não for nulo
			// adiciona seu valor à soma e avança para o próximo nó de l1
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil { // Se o nó atual de l2 não for nulo
			// adiciona seu valor à soma e avança para o próximo nó de l2
			sum += l2.Val
			l2 = l2.Next
		}
		// Calcula o dígito atual e o transporte
		digit := sum % 10
		carry = sum / 10
		// Cria um novo nó com o dígito atual e o adiciona à nova lista
		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return dummy.Next // Retorna o próximo nó do nó dummy, que é o primeiro nó da nova lista
}
