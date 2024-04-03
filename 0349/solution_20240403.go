package _349

func intersection(nums1 []int, nums2 []int) []int {

	numsInA := make(map[int]int)
	for _, v := range nums1 {
		numsInA[v] += 1
	}

	interOfAB := make(map[int]int)
	for _, v := range nums2 {
		if numsInA[v] > 0 {
			interOfAB[v] = 1
		}
	}

	result := make([]int, 0, len(interOfAB))
	for v := range interOfAB {
		result = append(result, v)
	}
	return result
}
