package _167

func twoSum__(numbers []int, target int) []int {
	for i, one := range numbers {
		other := target - one
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := low + (high-low)/2
			if numbers[mid] < other {
				low = mid + 1
			} else if numbers[mid] > other {
				high = mid - 1
			} else {
				return []int{i + 1, mid + 1}
			}
		}
	}
	return []int{}
}
