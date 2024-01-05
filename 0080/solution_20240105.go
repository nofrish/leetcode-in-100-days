package _080

func removeDuplicates(nums []int) int {
	slow, fast, count := 0, 1, 1
	for ; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
			count = 1
		} else {
			if count == 1 {
				slow++
				nums[slow] = nums[fast]
				count = 2
			}
		}
	}
	return slow + 1
}

func removeDuplicates_(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}

	slow, fast := 2, 2
	for ; fast < len(nums); fast++ {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
