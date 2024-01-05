package _076

func minWindow(s string, t string) string {

	target := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		target[t[i]] += 1
	}

	window := make(map[uint8]int)
	slow, fast, result := 0, -1, ""

	for fast++; fast < len(s); fast++ {
		if _, ok := target[s[fast]]; !ok {
			continue
		}
		window[s[fast]] += 1
		if !satisfied(window, target) {
			continue
		}

		for satisfied(window, target) {
			if result == "" || len(result) > fast-slow+1 {
				result = s[slow : fast+1]
			}
			window[s[slow]] -= 1
			slow++
		}
	}

	return result
}

func satisfied(source, target map[uint8]int) bool {
	for key, v1 := range target {
		if v2 := source[key]; v2 < v1 {
			return false
		}
	}
	return true
}
