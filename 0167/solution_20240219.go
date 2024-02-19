package _167

func twoSum_(nums []int, target int) []int {
	for i, v := range nums {
		want := target - v
		low, high := i+1, len(nums)-1
		for low <= high {
			mid := low + (high-low)/2
			if nums[mid] < want {
				low = mid + 1
			} else if nums[mid] > want {
				high = mid - 1
			} else {
				return []int{i + 1, mid + 1}
			}
		}
	}
	return []int{}
}
