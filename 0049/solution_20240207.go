package _049

import "strings"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		chars := [26]int{}
		for i := 0; i < len(str); i++ {
			chars[int(str[i]-'a')] += 1
		}
		m[chars] = append(m[chars], str)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func groupAnagrams_(strs []string) [][]string {

	m := make(map[string][]string)
	for _, str := range strs {
		chars := [26]int{}
		for i := 0; i < len(str); i++ {
			chars[int(str[i]-'a')] += 1
		}
		sb := strings.Builder{}
		for i := 0; i < 26; i++ {
			for j := 0; j < chars[i]; j++ {
				sb.WriteByte('a' + byte(i))
			}
		}
		key := sb.String()

		if v, ok := m[key]; ok {
			m[key] = append(v, str)
		} else {
			m[key] = []string{str}
		}
	}

	result := make([][]string, 0)
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
