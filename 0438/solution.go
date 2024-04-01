package _438

func _findAnagrams(s string, p string) []int {

	if len(s) < len(p) {
		return []int{}
	}

	result := make([]int, 0)

	var charsOfP, chars [26]int
	for i := 0; i < len(p); i++ {
		charsOfP[p[i]-'a'] += 1
		chars[s[i]-'a'] += 1
	}

	if chars == charsOfP { // 初始状态判定
		result = append(result, 0)
	}
	for i := len(p); i < len(s); i++ { // 循环判定之后的状态
		chars[s[i]-'a'] += 1
		chars[s[i-len(p)]-'a'] -= 1
		if chars == charsOfP {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}
