package _014

func longestCommonPrefix_(strs []string) string {

	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:minLen]
}
