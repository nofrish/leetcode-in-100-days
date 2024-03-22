package _076

func minWindow_(s string, t string) string {

	satisfied := func(src, tgt map[byte]int) bool {
		for k, v := range tgt {
			if src[k] < v {
				return false
			}
		}
		return true
	}

	tgt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tgt[t[i]] += 1
	}

	slow, fast, result, exists := 0, -1, "", make(map[byte]int)
	for fast++; fast < len(s); fast++ {
		exists[s[fast]] += 1
		for ; satisfied(exists, tgt); slow++ {
			if result == "" || len(result) > fast-slow+1 {
				result = s[slow : fast+1]
			}
			exists[s[slow]] -= 1
		}
	}
	return result
}
