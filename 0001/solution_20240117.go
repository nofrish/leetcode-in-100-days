package _001

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, 0)
	for i, one := range nums {
		theOther := target - one
		if j, ok := m[theOther]; ok {
			return []int{i, j}
		}
		m[one] = i
	}

	return []int{}
}
