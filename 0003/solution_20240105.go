package _003

func lengthOfLongestSubstring(s string) int {

	m := make(map[uint8]int)

	slow, fast, max := 0, 0, 0
	for ; fast < len(s); fast++ {
		if idx, ok := m[s[fast]]; ok && idx >= slow {
			slow = idx + 1
		} else {
			if max < fast-slow+1 {
				max = fast - slow + 1
			}
		}
		m[s[fast]] = fast
	}

	return max
}
