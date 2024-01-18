package _006

import "strings"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	sbs := make([]strings.Builder, numRows)
	sbs[0].WriteByte(s[0])

	idx, step := 1, 1
	for i := 1; i < len(s); i++ {
		sbs[idx].WriteByte(s[i])
		if idx == numRows-1 {
			step = -1
		}
		if idx == 0 {
			step = 1
		}
		idx += step
	}

	result := strings.Builder{}
	for _, sb := range sbs {
		result.WriteString(sb.String())
	}
	return result.String()
}
