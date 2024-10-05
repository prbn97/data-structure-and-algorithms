// https://leetcode.com/problems/longest-common-prefix/
package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

}
func longestCommonPrefix(strs []string) string {

	var letters [][]string
	// letters := make([][]string, len(strs))
	for _, word := range strs {
		lettersOfWord := strings.Split(word, "")
		letters = append(letters, lettersOfWord)
	}

	var commonPrefix []string

	for i := range letters[0] { //iterate in the first word
		char := letters[0][i]
		common := true
		for j := 1; j < len(letters); j++ {
			if i >= len(letters[j]) || letters[j][i] != char {
				common = false
				break
			}
		}
		if !common {
			break
		}
		commonPrefix = append(commonPrefix, char)
	}

	output := strings.Join(commonPrefix, "")
	return output
}
