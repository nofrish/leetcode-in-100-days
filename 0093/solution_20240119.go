package _093

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses_(s string) (result []string) {

	var dfs func(s string, prefix []string)
	dfs = func(s string, prefix []string) {
		options := make([]string, 0)
		for i := 0; i < len(s) && i < 3; i++ {
			options = append(options, s[:i+1])
		}

		for _, option := range options {
			if len(s)-len(option) > (4-len(prefix))*3 {
				continue
			}
			if len(option) > 1 && option[0] == '0' {
				continue
			}
			if len(prefix) == 3 && len(option) != len(s) {
				continue
			}
			number, _ := strconv.Atoi(option)
			if number > 255 {
				continue
			}

			if len(prefix) == 3 {
				result = append(result, fmt.Sprintf("%s.%s.%s.%s", prefix[0], prefix[1], prefix[2], option))
			} else {
				dfs(s[len(option):], append(append([]string{}, prefix...), option))
			}
		}
	}

	dfs(s, []string{})
	return
}
