package _169

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v] += 1
		if counts[v] > len(nums)/2 {
			return v
		}
	}
	return -1
}

func majorityElement_(nums []int) int {

	votes, x := 0, 0
	for i := 0; i < len(nums); i++ {
		if votes == 0 {
			x = nums[i]
		}
		if nums[i] == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}
