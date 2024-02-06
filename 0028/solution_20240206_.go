package _028

func strStr_(s string, p string) int {

	result := make([]int, 0)

	next := make([]int, len(p)+1)
	next[0], next[1] = 0, 0
	for i, j := 0, 2; j <= len(p); j++ {
		for ; i > 0 && p[i] != p[j-1]; i = next[i] {
		}
		if p[i] == p[j-1] {
			i++
		}
		next[j] = i
	}

	for i, j := 0, 0; i < len(s); i++ {
		for ; j > 0 && p[j] != s[i]; j = next[j] {
		}
		if s[i] == p[j] {
			j++
		}
		if j == len(p) {
			result = append(result, i-j+1)
			j = next[j]
		}
	}

	if len(result) > 0 {
		return result[0]
	}
	return -1
}
