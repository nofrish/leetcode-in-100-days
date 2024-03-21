package _167

func twoSum___(nums []int, target int) []int {
	for i, num := range nums {
		other := target - num
		low, high := i+1, len(nums)-1
		for low <= high {
			mid := low + (high-low)/2
			if nums[mid] < other {
				low = mid + 1
			} else if nums[mid] > other {
				high = mid - 1
			} else {
				return []int{i + 1, mid + 1}
			}
		}
	}
	return []int{}
}
