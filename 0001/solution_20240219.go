package _001

func twoSum___(nums []int, target int) []int {
	exists := make(map[int]int)
	for i, v := range nums {
		want := target - v
		if j, ok := exists[want]; ok {
			return []int{i, j}
		}
		exists[v] = i
	}
	return []int{}
}
