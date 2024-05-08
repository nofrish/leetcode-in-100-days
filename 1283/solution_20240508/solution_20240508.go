package solution_20240508

func smallestDivisor(nums []int, threshold int) int {

	lo, hi := 1, 1
	for _, num := range nums {
		hi = max(num, hi)
	}

	for lo < hi {
		mid := lo + (hi-lo)/2
		if f(nums, mid) > threshold {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi
}

// 输入参数x为除数，输出为按要求计算的结果之和
// 可以知道，当x越大，则结果之和越小，当x越小，则结果之和越大
// 我们需要找到单调递减函数上满足结果的最小值
func f(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += (num + x - 1) / x
	}
	return sum
}
