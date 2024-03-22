package _076

func minWindow__(s string, t string) string {

	// window 表示窗口内的各个字符数
	// target 表示目标字符串的各个字符数
	window, target := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		target[t[i]] += 1
	}

	// satisfied 确认 window 能否满足 target
	satisfied := func() bool {
		for k, c := range target {
			if window[k] < c {
				return false
			}
		}
		return true
	}

	// 标准的最小窗口模版
	slow, fast, minStr := 0, -1, ""
	for fast++; fast < len(s); fast++ {
		window[s[fast]] += 1
		for ; satisfied(); slow++ {
			if minStr == "" || len(minStr) > fast-slow+1 {
				minStr = s[slow : fast+1]
			}
			window[s[slow]] -= 1
		}
	}

	return minStr
}
