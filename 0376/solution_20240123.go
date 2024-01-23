package _376

func wiggleMaxLength(nums []int) int {

	cur := 0
	for ; cur+1 < len(nums); cur++ {
		if nums[cur] != nums[cur+1] {
			break
		}
	}

	maxLen := 1
	for i := cur + 1; i < len(nums); i++ {
		if nums[i]-nums[cur] > 0 {
			// 沿着递增序列一直找
			for ; i+1 < len(nums); i++ {
				if nums[i+1] < nums[i] {
					cur = i
					break
				}
			}
		} else {
			// 沿着递减序列一直找
			for ; i+1 < len(nums); i++ {
				if nums[i+1] > nums[i] {
					cur = i
					break
				}
			}
		}
		maxLen++
	}

	return maxLen
}
