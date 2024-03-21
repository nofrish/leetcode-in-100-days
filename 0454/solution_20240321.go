package _454

func fourSumCount__(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	countAPlusB := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			countAPlusB[a+b] += 1
		}
	}

	result := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			result += countAPlusB[0-c-d]
		}
	}

	return result
}
