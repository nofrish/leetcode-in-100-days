package _395

func longestSubstring__(s string, k int) int {

	res := 0
	for distinct := 1; distinct <= 26; distinct++ {
		slow, fast, counts, valid := 0, -1, make(map[byte]int), 0
		for fast++; fast < len(s); fast++ {
			counts[s[fast]] += 1
			if counts[s[fast]] == k {
				valid += 1
			}
			for ; len(counts) > distinct; slow++ {
				counts[s[slow]] -= 1
				if counts[s[slow]] == k-1 {
					valid -= 1
				}
				if counts[s[slow]] == 0 {
					delete(counts, s[slow])
				}
			}
			if valid == distinct && res < fast-slow+1 {
				res = fast - slow + 1
			}
		}
	}
	return res
}
