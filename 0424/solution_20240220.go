package _424

func characterReplacement(s string, k int) int {
	slow, fast, counts, maxC, res := 0, -1, make(map[byte]int), 0, 0
	for fast++; fast < len(s); fast++ {
		counts[s[fast]] += 1
		maxC = max(maxC, counts[s[fast]])
		for ; fast-slow+1-maxC > k; slow++ {
			counts[s[slow]] -= 1
		}
		res = max(res, fast-slow+1)
	}
	return res
}

// 说明：
//   1. maxC 并不是当前窗口内出现次数最多字母的次数，而是历史窗口中出现次数最多字母的次数
//   2. 因为找到一个满足条件的窗口之后，接下来只需要去找一个更大的窗口, 要找更大的那个窗口，也就等价于找能够使得 maxC 增大的新的窗口
//   3. 也就是说在改变左窗口端点之后，判断新加入的字母的个数是否能超过 maxC(因为只有新加入来的字母才有可能使得最终的 maxC 增大)
