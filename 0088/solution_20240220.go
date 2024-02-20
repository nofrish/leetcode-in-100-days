package _088

func merge__(nums1 []int, m int, nums2 []int, n int) {

	i, j, cur := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; cur-- {
		if nums1[i] >= nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else {
			nums1[cur] = nums2[j]
			j--
		}
	}
	for ; j >= 0; j, cur = j-1, cur-1 {
		nums1[cur] = nums2[j]
	}
	return
}
