package _325

func maxSubArrayLen(nums []int, k int) (result int) {

	prefix := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}

	index := make(map[int]int)
	for i := 0; i < len(prefix); i++ {
		cur := prefix[i]
		expect := cur - k
		if j, ok := index[expect]; ok {
			result = max(result, i-j)
		}
		if _, ok := index[cur]; !ok {
			index[cur] = i
		}
	}

	return
}
