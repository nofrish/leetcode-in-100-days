package _395

import (
	"strings"
)

func longestSubstring(s string, k int) int {

	counts := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		counts[s[i]] += 1
	}

	var split byte
	for c, v := range counts {
		if v < k {
			split = c
			break
		}
	}
	if split == 0 {
		return len(s)
	}

	result := 0
	for _, subStr := range strings.Split(s, string(split)) {
		result = max(result, longestSubstring(subStr, k))
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubstring_(s string, k int) int {

	max := 0
	for i := 1; i <= 26; i++ {
		exists := make(map[byte]int, 0)
		slow, fast := 0, -1
		for fast++; fast < len(s); fast++ {
			exists[s[fast]] += 1
			for ; len(exists) > i; slow++ {
				exists[s[slow]] -= 1
				if exists[s[slow]] == 0 {
					delete(exists, s[slow])
				}
			}

			satisfied := true
			for _, v := range exists {
				if v < k {
					satisfied = false
					break
				}
			}
			if satisfied && max < fast-slow+1 {
				max = fast - slow + 1
			}
		}
	}

	return max
}
