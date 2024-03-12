package _713

func numSubarrayProductLessThanK(nums []int, k int) (result int) {
	if k <= 1 {
		return 0
	}
	slow, fast, prod := 0, -1, 1
	for fast++; fast < len(nums); fast++ {
		prod *= nums[fast]
		for ; prod >= k; slow++ {
			prod /= nums[slow]
		}
		result += fast - slow + 1
	}
	return
}
