package _266

func canPermutePalindrome(s string) bool {

	counts := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		counts[s[i]] += 1
	}

	odds := 0
	for _, v := range counts {
		if v%2 == 1 {
			odds++
			if odds > 1 {
				return false
			}
		}
	}

	return true
}
