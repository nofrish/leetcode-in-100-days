package _248

func numberOfSubarrays_(nums []int, k int) (result int) {

	// 举例：nums = [2 2 2 1 2 2 1 2 2 2 1 2], k = 2
	// 1. 当fast指针未指向第2个1时，即 len(next) < k 时，应该一直移动fast指针
	// 2. 当fast指针指向第2个1时，即 len(next) == k 时，slow 和 fast 之间有多个优美子数组，可以想象从slow到第一个1之间都是
	// 3. 当fast继续向前移动，都会有 nexts[0] - slow + 1 个 优美子数组出现
	// 4. 当fast移动到第3个1时，应该将slow直接移动到第一个1的后面，因此需要nexts数组来存储1的位置

	slow, fast, nexts := 0, -1, make([]int, 0)
	for fast++; fast < len(nums); fast++ {
		if nums[fast]%2 == 1 {
			nexts = append(nexts, fast)
		}
		if len(nexts) < k {
			continue
		}
		if len(nexts) > k {
			slow = nexts[0] + 1
			nexts = nexts[1:]
		}
		result += (nexts[0] - slow + 1)
	}
	return
}
