package _001

func twoSum____(nums []int, target int) []int {
	exists := make(map[int]int)
	for i, one := range nums {
		other := target - one
		if j, ok := exists[other]; ok {
			return []int{i, j}
		}
		exists[one] = i
	}
	return []int{}
}
