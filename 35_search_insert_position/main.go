// https://leetcode.com/problems/search-insert-position/
package main

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	searchInsert(nums, target)

}

func searchInsert(nums []int, target int) int {
	var targetWoldBeIndex int

	for i, num := range nums {
		if target == num {
			return i
		}
		if target < num {
			targetWoldBeIndex = i
			break
		}
		targetWoldBeIndex = i + 1
	}

	return targetWoldBeIndex
}
