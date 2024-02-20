package _003

func lengthOfLongestSubstring_______(s string) int {

	slow, fast, exists, res := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		exists[s[fast]] += 1
		for exists[s[fast]] > 1 {
			exists[s[slow]] -= 1
			slow++
		}
		if res < fast-slow+1 {
			res = fast - slow + 1
		}
	}
	return res
}
