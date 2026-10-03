func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right {
			if isAlphanumeric(rune(s[left])) {
				break
			} else {
				left++
			}
		}
		for left < right {
			if isAlphanumeric(rune(s[right])) {
				break
			} else {
				right--
			}
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
	       (c >= 'A' && c <= 'Z') ||
		   (c >= '0' && c <= '9')
}