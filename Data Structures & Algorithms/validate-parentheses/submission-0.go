func isValid(s string) bool {
    if len(s) % 2 != 0 { return false }

	stack := []rune{}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')':
			if len(stack) == 0 || stack[len(stack) - 1] != '(' {
				return false
			}
			stack = stack[:len(stack) - 1]
		case ']':
			if len(stack) == 0 || stack[len(stack) - 1] != '[' {
				return false
			}
			stack = stack[:len(stack) - 1]
		case '}':
			if len(stack) == 0 || stack[len(stack) - 1] != '{' {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
