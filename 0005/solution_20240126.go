package _005

func longestPalindrome_(s string) (result string) {

	getPString := func(s string, l, r int) string {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l, r = l-1, r+1
		}
		return s[l+1 : r]
	}

	for i := 0; i < len(s); i++ {
		p := getPString(s, i, i+1)
		if len(p) > len(result) {
			result = p
		}
		p = getPString(s, i-1, i+1)
		if len(p) > len(result) {
			result = p
		}
	}

	return
}
