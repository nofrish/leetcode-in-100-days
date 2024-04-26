package _153

// 解法一
// 在处理完没有旋转的情况之后，再通过 first 来找最小值
func _findMin_Solution1(nums []int) int {

	// 先处理没有旋转的情况
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	low, high, first := 0, len(nums)-1, nums[0]
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] < first {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low]
}

// 解法二
// 在处理完没有旋转的情况之后，再通过 last 来找最小值
func _findMin_Solution2(nums []int) int {

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	low, high, last := 0, len(nums)-1, nums[len(nums)-1]
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > last {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

// 解法三解法四同理，可以通过 first 和 last 来找最大值，然后通过最大值来找最小值
