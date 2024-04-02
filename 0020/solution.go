package _020

func _isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	pairs := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
