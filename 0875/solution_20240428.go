package _875

func minEatingSpeed_(piles []int, h int) int {

	// f(x) 输入吃香蕉的速度，返回需要的耗时
	// 显然，珂珂吃香蕉的速度越快，耗时越短
	// f(x) 是一个单调递减函数
	f := func(piles []int, k int) int {
		cost := 0
		for _, pile := range piles {
			cost += (pile + k - 1) / k
		}
		return cost
	}

	low, high := 1, 1
	for _, pile := range piles {
		high = max(high, pile)
	}

	for low < high {
		mid := (low + high) / 2
		if f(piles, mid) > h {
			low = mid + 1
		} else {
			high = mid // 让high锁定在可以满足条件的位置
		}
	}
	return high
}
