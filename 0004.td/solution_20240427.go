package _004_td

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1 // ensure len(nums1) >= len(nums2)
	}

}
