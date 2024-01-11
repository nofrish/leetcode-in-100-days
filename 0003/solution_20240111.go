package _003

import "math"

func lengthOfLongestSubstring___(s string) int {

	slow, fast, exists, max := 0, -1, make(map[byte]bool), 0

	for fast++; fast < len(s); fast++ {
		c := s[fast]
		if !exists[c] {
			exists[c] = true
			if max < fast-slow+1 {
				max = fast - slow + 1
			}
			continue
		}

		for ; exists[c]; slow++ {
			exists[s[slow]] = false
		}

		exists[c] = true
	}

	return max
}

func lengthOfLongestSubstring____(s string) int {

	slow, fast, exists, max := 0, -1, make(map[byte]bool), 0

	for fast++; fast < len(s); fast++ {
		c := s[fast]
		for ; exists[c]; slow++ {
			exists[s[slow]] = false
		}
		exists[c] = true

		max = int(math.Max(float64(max), float64(fast-slow+1)))
	}

	return max
}
