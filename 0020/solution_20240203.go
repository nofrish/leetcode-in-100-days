package _020

func isValid(s string) bool {

	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}

		switch s[i] {
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
		}
		stack = stack[:len(stack)-1]
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
