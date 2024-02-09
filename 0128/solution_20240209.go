package _128

func longestConsecutive(nums []int) (result int) {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		if m[n] > 0 {
			continue
		}
		left := m[n-1]
		right := m[n+1]

		total := 1 + left + right
		if total > result {
			result = total
		}

		m[n] = total
		m[n-left] = total
		m[n+right] = total
	}
	return
}
