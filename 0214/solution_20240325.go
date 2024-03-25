package _214

import (
	"strings"
)

func shortestPalindrome_(s string) string {

	reverse := func(s string) string {
		sb := strings.Builder{}
		for i := len(s) - 1; i >= 0; i-- {
			sb.WriteByte(s[i])
		}
		return sb.String()
	}

	isPString := func(s string) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}

	for r := len(s); r > 0; r-- {
		if isPString(s[:r]) {
			return reverse(s[r:]) + s
		}
	}
	return ""
}
