package _003

import "math"

// Solution One
func _lengthOfLongestSubstring(s string) int {

	// 初始化的条件：注意此时fast指向-1
	// 因为如果fast指向0，那么[0,0]中实际上是有一个元素的，那这时的exists就和两个指针的状态不一致了，当然max也不一致
	slow, fast, exists, max := 0, -1, make(map[byte]bool), 0

	// 开始循环，fast++，相当于是开始检查第1个元素
	for fast++; fast < len(s); fast++ {
		c := s[fast]              // 对于每一个元素
		for ; exists[c]; slow++ { // 判断是否可以放入，如果无法放入，就移动slow，直到可以放入为止
			exists[s[slow]] = false
		}
		exists[c] = true // 放入该元素

		// 结果计算
		max = int(math.Max(float64(max), float64(fast-slow+1)))
	}

	return max
}

// Solution Two
func __lengthOfLongestSubstring(s string) int {
	// 初始化的条件：注意此时fast指向-1
	// 因为如果fast指向0，那么[0,0]中实际上是有一个元素的，那这时的exists就和两个指针的状态不一致了，当然max也不一致
	// 注意这里的exists存储的是每个元素的数量
	slow, fast, exists, max := 0, -1, make(map[byte]int), 0

	// 开始循环，fast++，相当于是开始检查第1个元素
	for fast++; fast < len(s); fast++ {
		// 先将元素放入并将数量加1，再判断数量是否大于1，持续移动slow，直到数量等于1
		for exists[s[fast]] += 1; exists[s[fast]] > 1; slow++ {
			exists[s[slow]] -= 1
		}
		// 结果计算
		max = int(math.Max(float64(max), float64(fast-slow+1)))
	}
	return max
}
