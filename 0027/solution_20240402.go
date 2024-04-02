package _027

func removeElement__(nums []int, val int) int {
	slow, fast := -1, -1
	for fast++; fast < len(nums); fast++ {
		if nums[fast] != val {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
