package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") == false`)
	}
}
