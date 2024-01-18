package _022

func generateParenthesis(n int) (result []string) {

	var dfs func(prefix string, left, right int)
	dfs = func(prefix string, left, right int) {
		if right > left || left > n || right > n {
			return
		}
		if left == n && right == n {
			result = append(result, prefix)
		}
		dfs(prefix+"(", left+1, right)
		dfs(prefix+")", left, right+1)
	}

	dfs("", 0, 0)
	return
}
