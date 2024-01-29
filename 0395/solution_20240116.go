package _395

func longestSubstring(s string, k int) int {

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
