package _020

func isValid__(s string) bool {

	pair := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if s[i] != pair[top] {
			return false
		}
	}
	return true
}
