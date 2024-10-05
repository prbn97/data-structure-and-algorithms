// https://leetcode.com/problems/length-of-last-word/
package main

import "strings"

func main() {
	s := "   fly me   to   the moon  "
	lengthOfLastWord(s)

}

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	lastWord := words[len(words)-1]
	return len(lastWord)

}
