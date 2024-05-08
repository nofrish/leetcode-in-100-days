package _389

func findTheDifference(s string, t string) byte {

	exists1, exists2 := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s); i++ {
		exists1[s[i]] += 1
	}
	for i := 0; i < len(t); i++ {
		exists2[t[i]] += 1
	}

	for k, v := range exists2 {
		if exists1[k] < v {
			return k
		}
	}
	return 0
}
