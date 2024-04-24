package _493

func longestSubarray(nums []int) (result int) {

	slow, fast, count := 0, -1, 0
	for fast++; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			count++
		}
		for ; count > 1; slow++ {
			if nums[slow] == 0 {
				count--
			}
		}
		result = max(result, fast-slow) // omit 0
	}
	return
}
