package _080

func removeDuplicates___(nums []int) int {

	slow, fast, count := 0, 0, 1
	for fast++; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
			count = 1
		} else if count == 1 {
			slow++
			nums[slow] = nums[fast]
			count++
		}
	}
	return slow + 1
}
