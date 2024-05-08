package _875

import "math"

func minEatingSpeed(piles []int, h int) int {

	lo, hi := 1, math.MaxInt
	for lo < hi {
		mid := lo + (hi-lo)/2
		if f(piles, mid) > h { // 如果吃完香蕉的速度大于h，说明吃的速度不够快
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi
}

// 输入是珂珂吃香蕉的速度，输出是珂珂吃完所有香蕉的时间
// 显然，珂珂吃香蕉的速度越快，吃完所有香蕉的时间就越短，因此这是一个单调递减函数
// 而我们是想要找到一个单调递减函数上满足需求的最小值
func f(piles []int, x int) int {
	sum := 0
	for _, pile := range piles {
		sum += (pile + x - 1) / x
	}
	return sum
}
