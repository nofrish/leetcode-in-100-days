package _005

func longestPalindrome__(s string) string {

	getPString := func(s string, l, r int) string {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return s[l+1 : r]
	}

	result := ""
	for i := 0; i < len(s); i++ {
		p := getPString(s, i, i+1)
		if len(result) < len(p) {
			result = p
		}
		p = getPString(s, i-1, i+1)
		if len(result) < len(p) {
			result = p
		}
	}
	return result
}
