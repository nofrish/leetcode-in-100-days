package _003

func lengthOfLongestSubstring________(s string) int {

	slow, fast, exists, result := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		exists[s[fast]] += 1
		for ; exists[s[fast]] > 1; slow++ {
			exists[s[slow]] -= 1
		}
		result = max(result, fast-slow+1)
	}
	return result
}
