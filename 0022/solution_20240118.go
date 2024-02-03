package _022

func generateParenthesis(n int) (result []string) {

	var backtrack func(prefix string, left, right int)
	backtrack = func(prefix string, left, right int) {
		if right > left || left > n || right > n {
			return
		}
		if left == n && right == n {
			result = append(result, prefix)
		}
		backtrack(prefix+"(", left+1, right)
		backtrack(prefix+")", left, right+1)
	}

	backtrack("", 0, 0)
	return
}
