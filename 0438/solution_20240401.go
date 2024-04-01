package _438

func findAnagrams_(s string, p string) []int {

	target := [26]int{}
	for i := 0; i < len(p); i++ {
		target[p[i]-'a'] += 1
	}

	result := make([]int, 0)
	for i := 0; i <= len(s)-len(p); i++ {
		source := [26]int{}
		for j := i; j < i+len(p); j++ {
			source[s[j]-'a'] += 1
		}
		if source == target {
			result = append(result, i)
		}
	}
	return result
}
