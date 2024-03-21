package _001

func twoSum_____(nums []int, target int) []int {
	exists := make(map[int]int)
	for i, num := range nums {
		other := target - num
		if j, ok := exists[other]; ok {
			return []int{i, j}
		}
		exists[num] = i
	}
	return []int{}
}
