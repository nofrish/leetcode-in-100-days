package _093

import (
	"strconv"
	"strings"
)

func restoreIpAddresses__(s string) (result []string) {

	var dfs func(prefix []string, remains string)
	dfs = func(prefix []string, remains string) {
		if len(prefix) == 4 && len(remains) > 0 {
			return
		}
		if len(prefix) == 4 && len(remains) == 0 {
			result = append(result, strings.Join(prefix, "."))
			return
		}
		for i := 0; i < len(remains) && i < 3; i++ {
			option := remains[0 : i+1]
			if len(option) > 1 && option[0] == '0' {
				return
			}
			v, _ := strconv.Atoi(option)
			if v > 255 {
				return
			}
			dfs(append(prefix, option), remains[i+1:])
		}
	}

	dfs([]string{}, s)
	return
}
