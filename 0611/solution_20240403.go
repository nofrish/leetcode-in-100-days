package _611

import (
	"sort"
)

func triangleNumber(nums []int) int {

	sort.Ints(nums)

	result := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ { // 遍历第一个值和第二个值，二分查找求第三个值
			target := nums[i] + nums[j] // 三角形两边之和大于第三边
			low, high, mid, found := j+1, len(nums)-1, -1, false
			for low <= high {
				mid = low + (high-low)/2
				if nums[mid] >= target {
					high = mid - 1
				} else {
					if mid == len(nums)-1 || nums[mid+1] >= target {
						found = true
						break
					}
					low = mid + 1
				}
			}
			if found {
				result += mid - j
			}
		}
	}
	return result
}
