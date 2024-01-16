package _003

func lengthOfLongestSubstring_____(s string) int {
	slow, fast, exists, max := 0, -1, make(map[byte]int, 0), 0
	for fast = 0; fast < len(s); fast++ {
		exists[s[fast]] += 1
		for ; exists[s[fast]] > 1; slow++ {
			exists[s[slow]] -= 1
		}
		if max < fast-slow+1 {
			max = fast - slow + 1
		}
	}
	return max
}
