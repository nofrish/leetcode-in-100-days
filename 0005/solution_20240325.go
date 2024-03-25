package _005

func longestPalindrome___(s string) string {

	getPString := func(s string, i, j int) string {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
		}
		return s[i+1 : j]
	}

	result := ""
	for i := 0; i < len(s); i++ {
		p := getPString(s, i-1, i+1)
		if len(p) > len(result) {
			result = p
		}
		p = getPString(s, i, i+1)
		if len(p) > len(result) {
			result = p
		}
	}
	return result
}
