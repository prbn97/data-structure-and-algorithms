// https://leetcode.com/problems/plus-one/
package main

import "fmt"

func main() {
	digits1 := []int{1, 2, 3}
	digits2 := []int{9}
	digits3 := []int{9, 9}
	plusOne(digits1)
	plusOne(digits2)
	plusOne(digits3)

}

func plusOne(digits []int) []int {
	fmt.Println(digits)

	lastIndex := len(digits) - 1
	digits[lastIndex] += 1

	for i := lastIndex; i >= 0; i-- {
		if digits[i] == 10 { // if 10
			digits[i] = 0 // set 0

			if i == 0 { // if first number
				digits = append([]int{1}, digits...) //add to the left
			} else {
				digits[i-1] += 1 // Incrementa o dígito anterior
			}
		} else {
			break
		}

	}

	fmt.Println(digits)
	return digits

}
