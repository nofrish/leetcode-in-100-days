package _647

func countSubstrings(s string) int {

	count := 0
	countPString := func(s string, left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		countPString(s, i, i)
		countPString(s, i, i+1)
	}
	return count
}
