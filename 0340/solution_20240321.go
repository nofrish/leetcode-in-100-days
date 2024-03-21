package _340

func lengthOfLongestSubstringKDistinct___(s string, k int) int {
	slow, fast, exists, maxLen := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		exists[s[fast]] += 1
		for ; len(exists) > k; slow++ {
			exists[s[slow]] -= 1
			if exists[s[slow]] == 0 {
				delete(exists, s[slow])
			}
		}
		maxLen = max(maxLen, fast-slow+1)
	}
	return maxLen
}
