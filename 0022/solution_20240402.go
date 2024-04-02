package _022

func generateParenthesis__(n int) (result []string) {

	var dfs func(prefix []byte, left, right int)
	dfs = func(prefix []byte, left, right int) {
		if left > n || right > n || right > left {
			return // invalid cases
		}
		if left == n && right == n {
			result = append(result, string(prefix))
		}
		dfs(append(prefix, '('), left+1, right)
		dfs(append(prefix, ')'), left, right+1)
	}

	dfs([]byte{}, 0, 0)
	return
}
