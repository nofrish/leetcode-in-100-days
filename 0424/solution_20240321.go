package _424

// 一个分析：
// 本题显然要用滑动窗口来解决，那么需要存储滑动窗口的哪些信息呢？
// 关键是：滑动窗口中出现次数最多的字符出现的次数。
func characterReplacement_(s string, k int) int {

	getMaxChar := func(counts map[byte]int) int {
		maxC := 0
		for _, v := range counts {
			maxC = max(maxC, v)
		}
		return maxC
	}

	slow, fast, counts, maxC, maxLen := 0, -1, make(map[byte]int), 0, 0
	for fast++; fast < len(s); fast++ {
		counts[s[fast]] += 1
		maxC = max(maxC, counts[s[fast]])
		for ; getMaxChar(counts)+k < fast-slow+1; slow++ {
			counts[s[slow]] -= 1
		}
		maxLen = max(maxLen, fast-slow+1)
	}

	return maxLen
}
