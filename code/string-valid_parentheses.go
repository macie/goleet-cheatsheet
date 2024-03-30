package main

func validParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '{' {
			stack = append(stack, '}')
		} else if c == '[' {
			stack = append(stack, ']')
		} else if len(stack) == 0 ||
			stack[len(stack)-1] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
