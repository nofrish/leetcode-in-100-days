package _454

func fourSumCount_(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	merge1 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			merge1[nums1[i]+nums2[j]] += 1
		}
	}

	merge2 := make(map[int]int)
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			merge2[nums3[i]+nums4[j]] += 1
		}
	}

	result := 0
	for one, c1 := range merge1 {
		other := 0 - one
		if c2, ok := merge2[other]; ok {
			result += c1 * c2
		}
	}
	return result
}
