package _080

func removeDuplicates__(nums []int) int {
	slow, fast, count := 0, 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] == nums[slow] {
			if count < 2 {
				slow++
				nums[slow] = nums[fast]
				count++
			}
		} else {
			slow++
			nums[slow] = nums[fast]
			count = 1
		}
	}
	return slow + 1
}
