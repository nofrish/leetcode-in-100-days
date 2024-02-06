package _438

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	charsOfP := [26]int{}
	chars := [26]int{}
	for i := 0; i < len(p); i++ {
		charsOfP[int(p[i]-'a')] += 1
		chars[int(s[i])-'a'] += 1
	}

	result := make([]int, 0)
	for i := len(p) - 1; i < len(s); i++ {
		if i != len(p)-1 {
			chars[int(s[i]-'a')] += 1
			chars[int(s[i-len(p)]-'a')] -= 1
		}
		if chars == charsOfP {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}
