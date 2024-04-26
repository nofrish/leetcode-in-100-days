package _153

func findMin__(nums []int) int {

	// 基本思路：
	// 1. 应用二分法的基本前提是：数据具有二段性，而不是必须递增。递增的本质其实也是提供了二段性，即相对于目标值来说，一段大一段小。
	// 2. 利用low和high不断逼近左边那一段的最大值，最后在最大值的位置相遇。因此，循环条件写为low<high，当low==high的时候，退出。
	// 3. 我们需要找一个值来作为二段性的参照，而我们应该选择nums[0]，这样可以兼顾旋转和未旋转两种情况。
	// 4. 由于low的变化是low=mid，因此需要让mid偏向high，所以需要 mid := low + high + 1 >> 1

	// 基本思路：让low和high不断逼近左边那一段的最大值，最后在最大值的位置相遇

	low, high := 0, len(nums)-1
	for low < high {
		mid := low + high + 1>>2
		if nums[mid] >= nums[0] {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return nums[(high+1)%len(nums)]
}
