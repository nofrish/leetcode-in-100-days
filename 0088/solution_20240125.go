package _088

func merge_(nums1 []int, m int, nums2 []int, n int) {
	i, j, c := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; c-- {
		if nums1[i] > nums2[j] {
			nums1[c] = nums1[i]
			i--
		} else {
			nums1[c] = nums2[j]
			j--
		}
	}
	for ; j >= 0; j, c = j-1, c-1 {
		nums1[c] = nums2[j]
	}
	return
}
