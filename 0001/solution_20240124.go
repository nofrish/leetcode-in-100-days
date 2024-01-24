package _001

func twoSum_(nums []int, target int) []int {
	exists := make(map[int]int, 0)
	for i, one := range nums {
		other := target - one
		if v, ok := exists[other]; ok {
			return []int{i, v}
		}
		exists[one] = i
	}

	return []int{}
}
