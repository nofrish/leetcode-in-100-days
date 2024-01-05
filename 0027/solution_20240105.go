package _027

func removeElement(nums []int, val int) int {
	slow, fast := -1, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
