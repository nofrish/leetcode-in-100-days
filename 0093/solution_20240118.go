package _093

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) (result []string) {

	var dfs func(s string, prefix []string)
	dfs = func(s string, prefix []string) {
		if (len(prefix) > 4) || (len(prefix) == 4 && len(s) > 0) {
			return
		}
		if len(prefix) == 4 && len(s) == 0 {
			result = append(result, strings.Join(prefix, "."))
		}
		for i := 0; i < len(s) && i < 3; i++ {
			guess := s[:i+1]
			if len(guess) > 1 && guess[0] == '0' {
				return
			}
			number, _ := strconv.Atoi(guess)
			if number > 255 {
				return
			}
			dfs(s[i+1:], append(prefix, guess))
		}
	}

	dfs(s, []string{})
	return
}
