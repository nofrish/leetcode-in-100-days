package _088

func merge(nums1 []int, m int, nums2 []int, n int) {

	p1, p2, px := m-1, n-1, m+n-1
	for ; p1 >= 0 && p2 >= 0; px-- {
		if nums1[p1] >= nums2[p2] {
			nums1[px] = nums1[p1]
			p1--
		} else {
			nums1[px] = nums2[p2]
			p2--
		}
	}

	for ; p2 >= 0; p2, px = p2-1, px-1 {
		nums1[px] = nums2[p2]
	}

	return
}
