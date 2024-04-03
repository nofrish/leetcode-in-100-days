package _162

func findPeakElement(nums []int) int {

	low, high := 0, len(nums)-1
	for low < high {
		// 如果循环条件是 low <= high ，那下面应该有一个退出条件是：
		// if low == high { return low }
		// 如果没有这个退出条件，循环会卡在 low == high == mid 上
		mid := low + (high-low)/2
		if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	// 下面也可以是 return high
	// 因为此时 low == high
	return low
}
