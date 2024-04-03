// https://leetcode.com/problems/two-sum/
// big o notation
package main

import "fmt"

func main() {

	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {

	result := []int{}

	for i, num1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			num2 := nums[j]

			if num1+num2 == target {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}
