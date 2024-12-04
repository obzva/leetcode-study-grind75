/*
Big O
- N: the length of the given string `s`
- Time complexity: O(N)
  - Since we have to iterate over `s`, the time complexity grows linearly
- Space complexity: O(N)
  - Stack could store up to N parentheses in worst case
*/

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)) // Stack to store parentheses
	for _, p := range s {
		if p == '(' || p == '{' || p == '[' { // Just append the parenthese to the stack
			stack = append(stack, p)
		} else {
			if len(stack) == 0 { // Invalid case; return false right away
				return false
			}
			last := stack[len(stack)-1]
			if (last == '(' && p == ')') || (last == '{' && p == '}') || (last == '[' && p == ']') { // Valid pair matched; pop the pair from the stack
				stack = stack[:len(stack)-1]
			} else { // Invalid case; return false right away
				return false
			}
		}
	}
	return len(stack) == 0 // Return true only if the stack is empty
}
