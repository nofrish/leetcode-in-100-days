package _209

func minSubArrayLen____(target int, nums []int) int {

	slow, fast, sum, minLen := 0, -1, 0, 0
	for fast++; fast < len(nums); fast++ {
		sum += nums[fast]
		for ; sum >= target; slow++ {
			if minLen == 0 || minLen > fast-slow+1 {
				minLen = fast - slow + 1
			}
			sum -= nums[slow]
		}
	}
	return minLen
}
