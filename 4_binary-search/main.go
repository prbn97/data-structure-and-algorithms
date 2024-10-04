// https://leetcode.com/problems/median-of-two-sorted-arrays/
// 4. Median of Two Sorted Arrays
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

// Constraints:
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

// Example:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

package main

import (
	"fmt"
	"math"
)

func main() {

	nums1, nums2 := []int{1, 3}, []int{2}    //2
	nums3, nums4 := []int{1, 2}, []int{3, 4} //2.5

	fmt.Println(nums1, nums2)
	fmt.Println(findMedianSortedArrays(nums1, nums2)) //2
	fmt.Println(nums3, nums4)
	fmt.Println(findMedianSortedArrays(nums3, nums4)) //2.5
}

// Algoritm: binary search - O(log (m+n).

// Median of Two Sorted and orded Arrays of diferrent sizes.
// x -> [x1, x2, x3, x4]
// y -> [y1, y2, y3, y4, y5, y6]
//
//
// Passo #1 (Aplicar a busca binaria no menor array)
// Dividir ambos arrays (x e y) usando duas partições (partitionX e partitionY) que serão indexes para dividir o array em lado esquerdo e direito.
// essa divisão em lados do array vai precisar respeitar duas condições.
//
// 1ª condição: o numero de elementos à esquerda da partição dos array igual ao numero de elementos a direita.
// obs.: caso o total de elemntos seja impar teremos um elemento a mais à esquerda.
//             Left X | Right X
//                    |
// array X            x[partitionX]
//      [  x1,   x2,  |  x3,  x4  ]
//                    |
//                    |
// array Y            y[partitionY]
//      [ y1, y2, y3, | y4, y5, y6  ]
//                    |
//                    |
//             Left Y | Right Y
//
// 2ª condição:
// Todos os elementos à esquerda de X são iguais ou menores a todos os elementos a direita de Y.
// Todos os elementos à esquerda de Y são iguais ou menores a todos os elementos a direita de X.
//  max(left_x, left_y) <= min(right_x, right_y)
//    maxLeft_X ≤ minRight_Y
//    maxLeft_Y ≤ minRight_X

// Passo #2 (Encontrar a mediana dos arrays combinados)
// Se o número total de elementos for ímpar --a mediana é o maior elemento na parte esquerda.
// max(maxLeft1, maxLeft2)
//
// Se o número total de elementos for par --a mediana é a média entre o maior elemento na parte esquerda e o menor elemento na parte direita.
// max(maxLeft_X, maxLeft_Y) + min(minRight_Y, minRight_Y)) / 2.0
//

func findMedianSortedArrays(x []int, y []int) float64 {
	if len(x) > len(y) {
		return findMedianSortedArrays(y, x)
	}
	start := 0
	end := len(x)

	// Passo #1: Dividir ambos arrays (x e y)
	for start <= end {

		//1ª condição

		partitionX := (start + end) / 2
		partitionY := (len(x)+len(y)+1)/2 - partitionX

		maxLeftX, maxLeftY := math.MinInt32, math.MinInt32
		minRightX, minRightY := math.MaxInt32, math.MaxInt32

		if partitionX > 0 {
			maxLeftX = x[partitionX-1]
		}
		if partitionX < len(x) {
			minRightX = x[partitionX]
		}
		if partitionY > 0 {
			maxLeftY = y[partitionY-1]
		}
		if partitionY < len(y) {
			minRightY = y[partitionY]
		}

		//2ª condição
		if maxLeftX <= minRightY && maxLeftY <= minRightX {

			// Passo #2

			// Se o número total de elementos é ímpar
			if (len(x)+len(y))%2 == 1 {
				return float64(max(maxLeftX, maxLeftY))
			} else {
				// Se o número total de elementos é par
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			}
		} else if maxLeftX > minRightY {
			// Mover para a esquerda em x
			end = partitionX - 1
		} else {
			// Mover para a direita em x
			start = partitionX + 1
		}
	}

	return 0.0 // Este caso nunca deve ser alcançado devido às garantias do problema
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
