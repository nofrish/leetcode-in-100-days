package _026

func removeDuplicates_(nums []int) int {
	slow, fast := 0, 1
	for ; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
