package _567

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	var charsOfS1, chars [26]int
	for i := 0; i < len(s1); i++ {
		charsOfS1[s1[i]-'a'] += 1
		chars[s2[i]-'a'] += 1
	}

	if chars == charsOfS1 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		chars[s2[i]-'a'] += 1
		chars[s2[i-len(s1)]-'a'] -= 1
		if chars == charsOfS1 {
			return true
		}
	}
	return false
}
