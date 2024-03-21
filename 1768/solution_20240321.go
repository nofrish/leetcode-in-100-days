package _768

import "strings"

func mergeAlternately_(word1 string, word2 string) string {

	sb := strings.Builder{}

	shortLen := min(len(word1), len(word2))
	for i := 0; i < shortLen; i++ {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}

	if len(word1) > shortLen {
		sb.WriteString(word1[shortLen:])
	}
	if len(word2) > shortLen {
		sb.WriteString(word2[shortLen:])
	}
	return sb.String()
}
