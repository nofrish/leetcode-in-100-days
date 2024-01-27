package _167

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		v := target - nums[i]
		low, high := i+1, len(nums)-1
		for low <= high {
			mid := low + (high-low)/2
			if nums[mid] < v {
				low = mid + 1
			} else if nums[mid] > v {
				high = mid - 1
			} else {
				return []int{i + 1, mid + 1}
			}
		}
	}

	return []int{}
}
