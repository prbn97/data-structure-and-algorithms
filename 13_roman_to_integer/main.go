// https://leetcode.com/problems/roman-to-integer/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))

}

func romanToInt(s string) int {
	nurmerals := 0
	romanValue := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	romanNumbers := strings.Split(s, "")

	for i := 0; i < len(romanNumbers); i++ {
		value := romanValue[romanNumbers[i]]

		if i+1 < len(romanNumbers) && romanValue[romanNumbers[i+1]] > value {
			nurmerals -= value
		} else {
			nurmerals += value
		}
	}
	return nurmerals
}
