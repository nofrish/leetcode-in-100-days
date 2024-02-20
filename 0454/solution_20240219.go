package _454

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (result int) {

	m := make(map[int]int)

	for i := range nums1 {
		for j := range nums2 {
			m[nums1[i]+nums2[j]] += 1
		}
	}

	for i := range nums3 {
		for j := range nums4 {
			other := 0 - nums3[i] - nums4[j]
			if m[other] > 0 {
				result += m[other]
			}
		}
	}

	return
}
