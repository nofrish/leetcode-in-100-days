package _340

import "math"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	slow, fast, exists, max := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		for exists[s[fast]] += 1; len(exists) > k; slow++ {
			exists[s[slow]] -= 1
			if exists[s[slow]] == 0 {
				delete(exists, s[slow])
			}
		}
		max = int(math.Max(float64(max), float64(fast-slow+1)))
	}
	return max
}
