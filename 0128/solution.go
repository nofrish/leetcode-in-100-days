package _128

func _longestConsecutive(nums []int) (result int) {

	exists := make(map[int]bool)
	for _, num := range nums {
		exists[num] = true // 先标记所有存在的数字，相当于在坐标轴上打点
	}

	for num := range exists {
		if !exists[num-1] { // 针对一段连续打点中的第一个点，看看后面还有几个点
			curLen := 1
			for i := 1; exists[num+i]; i++ {
				curLen++
			}
			if curLen > result {
				result = curLen
			}
		}
	}
	return
}
