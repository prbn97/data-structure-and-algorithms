// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums), nums)

}
func removeDuplicates(nums []int) int {
	writeIndex := 0 //primeira posicao a ser realocada
	lastWriteNumber := -1000

	for _, num := range nums {
		// se o "numero atual" não é igual ao "ultimo n escrito"
		// 	adiciono no array.
		// 	atualizo como ultimo numero escrito.
		// 	eu acrescento no index escrito
		if num != lastWriteNumber {
			nums[writeIndex] = num
			lastWriteNumber = num
			writeIndex++
		}
	}
	return writeIndex
}

// func removeDuplicates(nums []int) int {
// 	var expectedNums []int
// 	seen := make(map[int]bool)

// 	for _, num := range nums {
// 		if !seen[num] {
// 			expectedNums = append(expectedNums, num)
// 			seen[num] = true
// 		}
// 	}
// 	for i, num := range expectedNums {
// 		nums[i] = num
// 	}

// 	return len(expectedNums)
// }
