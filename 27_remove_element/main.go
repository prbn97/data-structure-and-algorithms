// https://leetcode.com/problems/remove-element/
package main

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	removeElement(nums, val)

}

func removeElement(nums []int, val int) int {
	writeIndex := 0
	removeInt := val

	for _, num := range nums {
		if num != removeInt {
			nums[writeIndex] = num
			writeIndex++
		}
	}
	return writeIndex
}
