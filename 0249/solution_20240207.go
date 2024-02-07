package _249

import "strings"

func groupStrings(ss []string) [][]string {

	m := make(map[string][]string)
	for _, s := range ss {
		if len(s) == 1 {
			m["a"] = append(m["a"], s)
			continue
		}

		key := strings.Builder{}
		for i := 1; i < len(s); i++ {
			key.WriteByte((s[i] - s[i-1] + 26) % 26)
			key.WriteByte(',')
		}
		m[key.String()] = append(m[key.String()], s)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
