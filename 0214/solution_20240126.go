package _214

import "strings"

func shortestPalindrome(s string) string {

	check := func(s string) bool {
		if s == "" {
			return true
		}
		i, j := 0, len(s)-1
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	for i := len(s); i > 0; i-- {
		if check(s[:i]) {
			return getReverseCopy(s[i:]) + s
		}
	}
	return ""
}

func getReverseCopy(s string) string {
	sb := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
