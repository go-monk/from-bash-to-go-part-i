package word

// IsPalindrome reports whether s reads the same forward and backward.
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
