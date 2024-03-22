package _248

func numberOfSubarrays(nums []int, k int) (result int) {

	slow, fast, count, nexts := 0, -1, 0, make([]int, 0)
	for fast++; fast < len(nums); fast++ {
		if nums[fast]%2 == 1 {
			count++
			nexts = append(nexts, fast)
		}
		for ; count > k; slow++ {
			if nums[slow]%2 == 1 {
				count--
				nexts = nexts[1:]
			}
		}
		if count == k {
			result += nexts[0] - slow + 1
		}
	}
	return
}
