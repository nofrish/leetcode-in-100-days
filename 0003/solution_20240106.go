package _003

func lengthOfLongestSubstring_(s string) int {
	slow, fast, exists, max := 0, -1, make(map[byte]bool), 0
	for fast++; fast < len(s); fast++ {
		if !exists[s[fast]] {
			exists[s[fast]] = true
			if max < fast-slow+1 {
				max = fast - slow + 1
			}
			continue
		}

		for ; exists[s[fast]]; slow++ {
			exists[s[slow]] = false
		}

		exists[s[fast]] = true
	}

	return max
}

func lengthOfLongestSubstring__(s string) int {
	slow, fast, exists, max := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		if idx, ok := exists[s[fast]]; !ok || idx < slow {
			exists[s[fast]] = fast
			if max < fast-slow+1 {
				max = fast - slow + 1
			}
			continue
		}

		slow = exists[s[fast]] + 1
		exists[s[fast]] = fast
	}

	return max
}
