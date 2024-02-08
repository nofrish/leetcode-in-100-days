package _768

import "strings"

func mergeAlternately(word1 string, word2 string) string {

	sb := strings.Builder{}

	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}
	if i < len(word1) {
		sb.WriteString(word1[i:])
	}
	if i < len(word2) {
		sb.WriteString(word2[i:])
	}

	return sb.String()
}
