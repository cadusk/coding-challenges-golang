package palindrome

import (
	"strconv"
)

func reverse(s string) string {
	rev := ""
	for _, r := range s {
		rev = string(r) + rev
	}
	return rev
}

func IsPalindrome(n int) bool {
	original := strconv.Itoa(n)
	return original == reverse(original)
}
