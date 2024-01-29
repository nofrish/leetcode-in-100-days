package _159

func lengthOfLongestSubstringTwoDistinct_(s string) int {

	exists := make(map[byte]int, 0)

	slow, fast, max := 0, -1, 0
	for fast++; fast < len(s); fast++ {
		exists[s[fast]] += 1
		for ; len(exists) > 2; slow++ {
			exists[s[slow]] -= 1
			if exists[s[slow]] == 0 {
				delete(exists, s[slow])
			}
		}
		if max < fast-slow+1 {
			max = fast - slow + 1
		}
	}

	return max
}
