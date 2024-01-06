package _807

import (
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	kmap := make(map[string]string)
	for _, pair := range knowledge {
		kmap[pair[0]] = pair[1]
	}

	sb := strings.Builder{}
	slow, fast, state := 0, -1, 0
	for fast++; fast < len(s); fast++ {
		if state == 0 /* looking for '(' */ {
			if s[fast] != '(' {
				continue
			}
			state = 1
			sb.WriteString(s[slow:fast])
			slow = fast + 1
		}
		if state == 1 /* looking for ')' */ {
			if s[fast] != ')' {
				continue
			}
			state = 0
			key := s[slow:fast]
			if val, ok := kmap[key]; ok {
				sb.WriteString(val)
			} else {
				sb.WriteString("?")
			}
			slow = fast + 1
		}
	}
	sb.WriteString(s[slow:fast])

	return sb.String()
}
