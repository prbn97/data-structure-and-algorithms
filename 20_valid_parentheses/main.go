// https://leetcode.com/problems/valid-parentheses/
package main

func main() {
	s := "(]"
	isValid(s)

}

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
func isValid(s string) bool {

	var st stack

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' { //if we found an opening bracket
			st.Push(char)
		} else { //is closing bracket
			top := st.Pop()
			if top == 'f' {
				return false
			}
			if (top == '(' && char == ')') || (top == '[' && char == ']') || (top == '{' && char == '}') {
				continue
			} else {
				return false
			}
		}
	}
	if len(st) == 0 {
		return true
	}
	return false
}

type stack []rune

func (s *stack) Push(item rune) {
	*s = append(*s, item)
}
func (s *stack) Pop() rune {
	if len(*s) == 0 {
		return 'f'
	}
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item

}
