package _159

import "math"

func lengthOfLongestSubstringTwoDistinct(s string) int {

	slow, fast, exists, max := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		for exists[s[fast]] += 1; len(exists) > 2; slow++ {
			exists[s[slow]] -= 1
			if exists[s[slow]] == 0 {
				delete(exists, s[slow])
			}
		}
		max = int(math.Max(float64(max), float64(fast-slow+1)))
	}
	return max
}
