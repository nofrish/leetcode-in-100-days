package _131

func partition(s string) (result [][]string) {

	var dfs func(s string, prefix []string)
	dfs = func(s string, prefix []string) {
		if len(s) == 0 {
			result = append(result, prefix)
		}
		for i := 0; i < len(s); i++ {
			ss := s[0 : i+1]
			if isPalindrome(ss) {
				dfs(s[i+1:], append(append([]string{}, prefix...), ss))
			}
		}
	}

	dfs(s, []string{})
	return
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
