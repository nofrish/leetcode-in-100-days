package _283

func moveZeroes(nums []int) {
	slow, fast := -1, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			slow++
			nums[slow] = nums[fast]
		}
	}
	for slow++; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}

func moveZeroes_(nums []int) {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
