package _001

func twoSum__(nums []int, target int) []int {
	exists := make(map[int]int, 0)
	for i, one := range nums {
		other := target - one
		if pos, ok := exists[other]; ok {
			return []int{pos, i}
		}
		exists[one] = i
	}

	return []int{}
}
