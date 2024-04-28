package _475

import "sort"

func findRadius(houses []int, heaters []int) int {

	sort.Ints(houses)
	sort.Ints(heaters)

	f := func(x int) bool {
		for i, j := 0, 0; i < len(houses); i++ {
			// 将 j 移动到可以满足 heaters[j]+x < houses[i]，此时可能在左边或右边
			for j < len(heaters) && heaters[j]+x < houses[i] {
				j++
			}

			// 如果在左边的话，下面的条件肯定满足
			// 如果在右边的花，下面的条件需要进一步满足 heaters[j]-x <= houses[i]
			// 所以下面的条件实际上可以简化为 if j < len(heaters) && heaters[j]-x <= houses[i]
			if j < len(heaters) && heaters[j]-x <= houses[i] && houses[i] <= heaters[j]+x {
				continue
			}
			return false
		}
		return true
	}

	low, high := 0, max(houses[len(houses)-1], heaters[len(heaters)-1])
	for low < high {
		mid := (low + high) / 2
		if f(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
