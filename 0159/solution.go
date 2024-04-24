package _159

func _lengthOfLongestSubstringTwoDistinct(s string) (result int) {
	slow, fast, exists := 0, -1, make(map[byte]int)
	for fast++; fast < len(s); fast++ {
		exists[s[fast]] += 1
		for ; len(exists) > 2; slow++ {
			exists[s[slow]] -= 1
			if exists[s[slow]] == 0 {
				delete(exists, s[slow])
			}
		}
		result = max(result, fast-slow+1)
	}
	return
}
