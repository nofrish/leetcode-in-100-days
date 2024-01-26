package _125

import "strings"

func isPalindrome(s string) bool {

	check := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			sb.WriteString(strings.ToLower(string(s[i])))
		}
	}

	return check(sb.String())
}
