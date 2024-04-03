// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
package main

import "strings"

func mian() {
	haystack := "sadbutsad"
	needle := "sad"

	strStr(haystack, needle)

}
func strStr(haystack string, needle string) int {
	// Input: haystack = "sadbutsad", needle = "sad"
	var firstLetterOccurrenceIndex int
	needleSlice := strings.Split(needle, "")
	haystackSlice := strings.Split(haystack, "")

	for i, letter := range haystackSlice { // in for loop of haystack
		if letter == needleSlice[0] { // check if the letter in haystack[i] is first in needle

			firstLetterOccurrenceIndex = i // if true - save this index of the letter
			// Check if the subsequent letters match
			match := true
			for j := 1; j < len(needleSlice); j++ {
				if i+j >= len(haystackSlice) || haystackSlice[i+j] != needleSlice[j] {
					match = false
					break
				}
			}

			if match {
				return firstLetterOccurrenceIndex
			}
		}

	}
	return -1
}
