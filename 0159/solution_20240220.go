package _159

func lengthOfLongestSubstringTwoDistinct__(s string) int {

	slow, fast, counts, res := 0, -1, make(map[byte]int), 0
	for fast++; fast < len(s); fast++ {
		counts[s[fast]] += 1
		for len(counts) > 2 {
			counts[s[slow]] -= 1
			if counts[s[slow]] == 0 {
				delete(counts, s[slow])
			}
			slow++
		}
		if res < fast-slow+1 {
			res = fast - slow + 1
		}
	}
	return res
}
