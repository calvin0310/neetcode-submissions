func isPalindrome(s string) bool {
	order := []rune{}

	for _, char := range s {
		if isAlphanumeric(char) {
			order = append(order, unicode.ToLower(char))
		} 
	}

	for i, o := range order {
		if o != order[len(order)-1-i] {
			return false
		}
	}

	return true
}

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		   (c >= 'A' && c <= 'Z') ||
		   (c >= '0' && c <= '9')
}