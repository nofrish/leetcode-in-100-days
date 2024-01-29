package _003

func lengthOfLongestSubstring______(s string) int {

	exists := make(map[byte]int, 0)

	slow, fast, max := 0, -1, 0
	for fast++; fast < len(s); fast++ {
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
