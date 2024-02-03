package _022

func generateParenthesis_(n int) (result []string) {

	var backtrack func(left, right int, prefix string)
	backtrack = func(left, right int, prefix string) {
		if right > left {
			return
		}
		if left == right && left == n {
			result = append(result, prefix)
		}
		if left < n {
			backtrack(left+1, right, prefix+"(")
		}
		if right < n {
			backtrack(left, right+1, prefix+")")
		}
	}

	backtrack(0, 0, "")
	return
}
