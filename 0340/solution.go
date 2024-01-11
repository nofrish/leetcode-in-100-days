package _340

import "math"

func _lengthOfLongestSubstringKDistinct(s string, k int) int {
	slow, fast, exists, max := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		// 先将新的值加入进来，再判断窗口的状态是否符合要求，如果不符合要求，则移动slow，直至符合要求
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
