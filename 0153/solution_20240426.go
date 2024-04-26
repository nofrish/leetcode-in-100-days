package _153

func findMin_(nums []int) int {

	// 先处理数组实际上没有旋转的情况
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	// 针对数组有旋转的情况，只和nums[0]进行对比
	// 其中low指针可以移动到第二段的第一个元素，high指针只能在第二段移动
	// 循环的退出条件是，low和high在最小的元素处相遇
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] >= nums[0] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}
