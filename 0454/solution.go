package _454

func _fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (result int) {

	countAB := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			countAB[a+b] += 1
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			result += countAB[-c-d]
		}
	}

	return
}
