package _272

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	result := make([][]int, 0)
	for _, interval := range intervals {
		if interval[1] < toBeRemoved[0] || interval[0] > toBeRemoved[1] {
			result = append(result, interval)
			continue
		}
		if interval[0] < toBeRemoved[0] {
			result = append(result, []int{interval[0], toBeRemoved[0]})
		}
		if interval[1] > toBeRemoved[1] {
			result = append(result, []int{toBeRemoved[1], interval[1]})
		}
	}
	return result
}
