package _005

func _longestPalindrome(s string) string {
	result := ""
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
	return result
}

func getPString(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
