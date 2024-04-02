package _810

import (
	"slices"
	"strings"
)

func finalString(s string) string {

	ss := [2][]byte{}
	dir := 1

	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			dir ^= 1
		} else {
			ss[dir] = append(ss[dir], s[i])
		}
	}
	slices.Reverse(ss[dir^1])
	return string(append(ss[dir^1], ss[dir]...))
}

func finalString_(s string) string {

	reverse := func(s string) string {
		sb := strings.Builder{}
		for i := len(s) - 1; i >= 0; i-- {
			sb.WriteByte(s[i])
		}
		return sb.String()
	}

	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			temp := sb.String()
			sb.Reset()
			sb.WriteString(reverse(temp))
		} else {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
